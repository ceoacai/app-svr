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
	"github.com/globalways/utils_go/errors"
	"github.com/globalways/utils_go/security"
	"net/http"
	"github.com/globalways/utils_go/convert"
	"regexp"
	"github.com/globalways/utils_go/consts"
)

type LoginController struct {
	BaseController
}

type ReqLogin struct {
	UserName string `form:"username"`
	PassWord string `form:"password"`
}

// login username is telphone?
func (req *ReqLogin) IsTel() bool {
	if b, err := regexp.MatchString(consts.Regexp_Mobile, req.UserName); !b || err != nil {
		return false
	}

	return true
}

func (req *ReqLogin) IsEmail() bool {
	return false
}

func (req *ReqLogin) IsHongId() bool {
	return false
}

type BodyLogin struct {
	Id       int64  `json:"id"`
	HongId   string `json:"hongid"`
	Tel      string `json:"tel"`
	Email    string `json:"email"`
	NickName string `json:"nickname"`
}

// curl -i -d "username=18610889275&password=123456" 127.0.0.1:8082/v1/hongid/login
// @router /login [post]
func (c *LoginController) Login() {
	reqLogin := new(ReqLogin)
	if err := c.ParseForm(reqLogin); err != nil {
		c.renderInternalError()
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
		rsp, err = c.forwardHttp("GET", fmt.Sprintf(hongIdInfoByTel, hongIdHost, reqLogin.UserName), nil)
	case reqLogin.IsEmail():
	case reqLogin.IsHongId():
	default:
		c.renderJson(errors.NewClientRsp(errors.CODE_BISS_ERR_USER_NAME))
		return
	}

	if err != nil {
		c.renderInternalError()
		return
	}
	defer rsp.Body.Close()

	clientRsp := new(errors.ClientRsp)
	if err := json.Unmarshal(c.getForwardHttpBody(rsp.Body), clientRsp); err != nil {
		c.renderInternalError()
		return
	}

	switch clientRsp.Status.Code {
	case errors.CODE_BISS_ERR_USER_NAME:
		c.renderJson(errors.NewClientRsp(errors.CODE_BISS_ERR_USER_NAME))
		return
	case errors.CODE_SUCCESS:
		break
	default:
		c.renderInternalError()
		return
	}

	body, ok := clientRsp.Body.(map[string]interface {})
	if !ok {
		c.renderInternalError()
		return
	}

	if !security.CompareHashAndPassword(body["PassWord"].(string), reqLogin.PassWord) {
		c.renderJson(errors.NewClientRsp(errors.CODE_BISS_ERR_PASSWORD))
		return
	}

	rspLogin := new(errors.ClientRsp)
	rspLogin.Status = errors.NewStatus(errors.CODE_SUCCESS)
	rspLogin.Body = &BodyLogin{
		Id:       convert.Float642Int64(body["Id"].(float64)),
		HongId:   body["HongId"].(string),
		Tel:      body["Tel"].(string),
		Email:    body["Email"].(string),
		NickName: body["NickName"].(string),
	}

	c.renderJson(rspLogin)
}
