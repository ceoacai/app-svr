// Copyright 2014 mit.zhao.chiu@gmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.
package controllers

import (
	"encoding/json"
	"fmt"
	hm "github.com/globalways/hongId_models/models"
	"github.com/globalways/utils_go/errors"
	"github.com/globalways/utils_go/security"
	"log"
	"net/http"
)

type LoginController struct {
	BaseController
}

type ReqLogin struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

// login username is telphone?
func (req *ReqLogin) IsTel() bool {
	return true
}

func (req *ReqLogin) IsEmail() bool {
	return false
}

func (req *ReqLogin) IsHongId() bool {
	return false
}

type RspLogin struct {
	Status *errors.StatusRsp `json:"status"`
	Body   *BodyLogin        `json:"body"`
}

type BodyLogin struct {
	Id       int64  `json:"id"`
	HongId   string `json:"hongid"`
	Tel      string `json:"tel"`
	Email    string `json:"email"`
	NickName string `json:"nickname"`
}

// curl -i -H "Content-Type: application/json" -d '{"username": "18610889275", "password": "123456"}' 127.0.0.1:8082/v1/hongid/login
// @router /login [post]
func (c *LoginController) Login() {
	reqLogin := new(ReqLogin)
	if err := json.Unmarshal(c.getHttpBody(), reqLogin); err != nil {
		c.renderJson(errors.NewErrorRsp(errors.CODE_SYS_ERR_BASE))
		return
	}

	// 验证输入参数
	c.validation(reqLogin)

	// handle http request param
	if c.handleParamError() {
		return
	}

	var rsp *http.Response
	var err error
	switch {
	case reqLogin.IsTel():
		log.Println("是手机号登录")
		rsp, err = c.forwardHttp("GET", fmt.Sprintf(hongIdInfoByTel, hongIdHost, reqLogin.UserName), nil)
	case reqLogin.IsEmail():
	case reqLogin.IsHongId():
	default:
		c.renderJson(errors.NewErrorRsp(errors.CODE_BISS_ERR_USER_NAME))
		return
	}

	if err != nil || rsp.StatusCode == http.StatusInternalServerError {
		c.renderInternalError()
		return
	}

	if rsp.StatusCode == http.StatusNotFound {
		c.renderJson(errors.NewErrorRsp(errors.CODE_BISS_ERR_USER_NAME))
		return
	}

	member := new(hm.Member)
	if err := json.Unmarshal(c.getForwardHttpBody(rsp.Body), member); err != nil {
		c.renderInternalError()
		return
	}

	if !security.CompareHashAndPassword(member.PassWord, reqLogin.PassWord) {
		c.renderJson(errors.NewErrorRsp(errors.CODE_BISS_ERR_PASSWORD))
		return
	}

	rspLogin := new(RspLogin)
	rspLogin.Status = errors.NewStatusRsp(errors.CODE_SUCCESS)
	body := &BodyLogin{
		Id:       member.Id,
		HongId:   member.HongId,
		Tel:      member.Tel,
		Email:    member.Email,
		NickName: member.NickName,
	}
	rspLogin.Body = body

	c.renderJson(rspLogin)
}
