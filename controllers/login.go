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
	"github.com/globalways/utils_go/errors"
	"net/http"
	hm "github.com/globalways/hongId_models/models"
	"github.com/globalways/utils_go/security"
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
	Id  int64 `json:"id"`
	UUID string `json:"uuid"`
	HongId string `json:"hongid"`
	NickName string `json:"nickname"`
}

// curl -i -H "Content-Type: application/json" -d '{"username": "18610889275", "password": "123456"}' 127.0.0.1:8082/v1/hongid/login
// @router /login [post]
func (c *LoginController) Login() {
	reqLogin := new(ReqLogin)
	if err := json.Unmarshal(c.getHttpBody(), reqLogin); err != nil {
		c.appenWrongParams(errors.NewFieldError("reqBody", err.Error()))
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
		rsp, err = c.forwardHttp("GET", hongIdHost+hongIdInfoByTel+reqLogin.UserName, nil)
	case reqLogin.IsEmail():
	case reqLogin.IsHongId():
	default:
		c.setHttpStatus(http.StatusForbidden)
		c.renderJson(errors.NewCommonOutRsp(errors.New(errors.CODE_BISS_ERR_USER_NAME)))
		return
	}

	if err != nil || rsp.StatusCode == http.StatusInternalServerError {
		c.renderInternalError()
		return
	}

	member := new(hm.Member)
	if err := json.Unmarshal(c.getForwardHttpBody(rsp.Body), member); err != nil {
		c.renderInternalError()
		return
	}

	if !security.CompareHashAndPassword(member.PassWord, reqLogin.PassWord) {
		c.setHttpStatus(http.StatusForbidden)
		c.renderJson(errors.NewCommonOutRsp(errors.New(errors.CODE_BISS_ERR_PASSWORD)))
		return
	}

	rspLogin := new(RspLogin)
	rspLogin.Id = member.Id
	rspLogin.UUID = member.UUID
	rspLogin.HongId = member.HongId
	rspLogin.NickName = member.NickName
	c.renderJson(rspLogin)
}


