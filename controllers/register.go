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
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/validation"
	hm "github.com/globalways/hongId_models/models"
	"github.com/globalways/utils_go/consts"
	"github.com/globalways/utils_go/errors"
	"github.com/globalways/utils_go/security"
	"regexp"
)

type RegisterController struct {
	BaseController
}

type MemberTelAtk struct {
	Tel  string `form:"tel"`
	Code string `form:"code"`
}

func (p *MemberTelAtk) Valid(v *validation.Validation) {
	if b, err := regexp.MatchString(consts.Regexp_Mobile, p.Tel); !b || err != nil {
		v.SetError("tel", "暂不支持此手机号码.")
	}

	if len(p.Code) == 0 {
		v.SetError("code", "手机验证码不能为空.")
	}
}

// app 手机注册
type ReqRegisterMemberByTel struct {
	Tel   string `json:"tel"`
	Group int64  `json:"group"`
}

// curl -i -d "tel=18610889275&code=123456" 123.57.132.7:8082/v1/hongid/register/smschk
// curl -i -d "tel=18610889275&code=123456" 127.0.0.1:8082/v1/hongid/register/smschk
// @router /register/smschk [post]
func (c *RegisterController) SmsCodeAtk() {
	// 解析httpbody
	memberTelAtk := new(MemberTelAtk)
	if err := c.ParseForm(memberTelAtk); err != nil {
		c.renderInternalError()
		return
	}

	// 验证输入参数
	c.validation(memberTelAtk)

	// handle http request param
	if c.handleParamError() {
		return
	}

	// 验证手机验证码正确性
	//	if !c.varifySmsAuthCode(memberTelAtk.Tel, memberTelAtk.Code) {
	//		c.renderJson(errors.NewClientRsp(errors.CODE_BISS_ERR_SMS_CODE))
	//		return
	//	}

	// 转发http请求
	reqMsg := &ReqRegisterMemberByTel{
		Tel:   memberTelAtk.Tel,
		Group: appUserGroupId,
	}
	reqBytes, err := json.Marshal(reqMsg)
	if err != nil {
		c.renderInternalError()
		return
	}
	rsp, err := c.forwardHttp("POST", fmt.Sprintf(hongIdRegByTel, hongIdHost), bytes.NewReader(reqBytes))
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
	case errors.CODE_HTTP_ERR_INVALID_PARAMS:
		c.renderJson(errors.NewClientRspf(errors.CODE_HTTP_ERR_INVALID_PARAMS, clientRsp.Status.Message))
		return
	case errors.CODE_BISS_ERR_REG:
		c.renderJson(errors.NewClientRsp(errors.CODE_BISS_ERR_REG))
		return
	case errors.CODE_BISS_ERR_TEL_ALREADY_IN:
		c.renderJson(errors.NewClientRsp(errors.CODE_BISS_ERR_TEL_ALREADY_IN))
		return
	case errors.CODE_SUCCESS:
		c.renderJson(errors.NewGlobalwaysErrorRsp(errors.ErrorOK()))
		return
	default:
		c.renderInternalError()
		return
	}
}

type MemberRegister struct {
	Tel      string `form:"tel"`
	NickName string `form:"nickname"`
	Password string `form:"password"`
}

func (p *MemberRegister) Valid(v *validation.Validation) {
	if b, err := regexp.MatchString(consts.Regexp_Mobile, p.Tel); !b || err != nil {
		v.SetError("tel", "暂不支持此手机号码.")
	}

	if len(p.NickName) == 0 {
		v.SetError("nickname", "昵称不能为空.")
	}

	if len(p.Password) == 0 {
		v.SetError("password", "密码不能为空.")
	}
}

// curl -i -d "tel=18610889275&nickname=mint&password=123456" 123.57.132.7:8082/v1/hongid/register/info
// curl -i -d "tel=18610889275&nickname=mint&password=123456" 127.0.0.1:8082/v1/hongid/register/info
// @router /register/info [post]
func (c *RegisterController) Register() {
	// 解析httpbody
	memberReg := new(MemberRegister)
	if err := c.ParseForm(memberReg); err != nil {
		c.renderInternalError()
		return
	}

	// 验证输入参数
	c.validation(memberReg)

	// handle http request param
	if c.handleParamError() {
		return
	}

	rspMemberInfo, err := c.forwardHttp("GET", fmt.Sprintf(hongIdInfoByTel, hongIdHost, memberReg.Tel), nil)
	if err != nil {
		c.renderInternalError()
		return
	}
	defer rspMemberInfo.Body.Close()

	clientRsp := new(errors.ClientRsp)
	if err := json.Unmarshal(c.getForwardHttpBody(rspMemberInfo.Body), clientRsp); err != nil {
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

	body, ok := clientRsp.Body.(map[string]interface{})
	if !ok {
		c.renderInternalError()
		return
	}
	args := make(map[string]interface{})
	args["NickName"] = memberReg.NickName
	args["PassWord"] = security.GenerateFromPassword(memberReg.Password)
	args["Status"] = hm.EMemberStatus_Use

	reqBytes, err := json.Marshal(&args)
	if err != nil {
		c.renderInternalError()
		return
	}

	rspMemberUpd, err := c.forwardHttp("PATCH", fmt.Sprintf(hongIdInfoById, hongIdHost, body["Id"]), bytes.NewReader(reqBytes))
	if err != nil {
		c.renderInternalError()
		return
	}
	defer rspMemberUpd.Body.Close()

	updRsp := new(errors.ClientRsp)
	if err := json.Unmarshal(c.getForwardHttpBody(rspMemberUpd.Body), updRsp); err != nil {
		c.renderInternalError()
		return
	}

	switch updRsp.Status.Code {
	case errors.CODE_HTTP_ERR_INVALID_PARAMS:
		c.renderJson(errors.NewClientRspf(errors.CODE_HTTP_ERR_INVALID_PARAMS, updRsp.Status.Message))
		return
	case errors.CODE_SUCCESS:
		c.renderJson(errors.NewGlobalwaysErrorRsp(errors.ErrorOK()))
		return
	default:
		c.renderInternalError()
		return
	}
}
