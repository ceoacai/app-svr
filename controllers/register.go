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
	hm "github.com/globalways/hongId_models/models"
	"github.com/globalways/utils_go/convert"
	"github.com/globalways/utils_go/errors"
	"github.com/globalways/utils_go/security"
	"net/http"
	"fmt"
)

type RegisterController struct {
	BaseController
}

/**
用户手机号注册流程
  1、app发送手机号到server [post]
  2、svrver发送请求到短信平台，发送短信，给app返回response信息 ok ? error
  3、app输入验证码，发送验证码&手机号到server [post]
  4、server验证用户输入与短信平台的验证码是否一致 ok ? error
  5、如果error，app显示错误信息，重新请求验证码 重复 1 - 4
  6、验证成功，app跳转下一界面，输入昵称 & 密码 & 重复密码，提交server [post]
  7、server完成注册
*/

type MemberTel struct {
	Tel string `valid:"Mobile" json:"tel"`
}

// curl -i -H "Content-Type: application/json" -i -d '{"tel":"18610889275"}' 123.57.132.7:8082/v1/hongid/register/smscode
// curl -i -H "Content-Type: application/json" -i -d '{"tel":"18610889275"}' 127.0.0.1:8082/v1/hongid/register/smscode
// @router /register/smscode [post]
func (c *RegisterController) SmsCode() {
	// 解析httpbody
	memberTel := new(MemberTel)
	if err := json.Unmarshal(c.getHttpBody(), memberTel); err != nil {
		c.appenWrongParams(errors.NewFieldError("memberTelPhone json", err.Error()))
	}

	// 验证手机号正确性
	c.validation(memberTel)

	// handle http request param
	if c.handleParamError() {
		return
	}

	// TODO 请求短信网关
	_, err := c.genSmsAuthCode(memberTel.Tel)
	if err != nil {
		c.renderJson(errors.NewCommonOutRsp(errors.New(errors.CODE_BISS_ERR_SMS_FAIL)))
		return
	}

	c.renderJson(errors.NewCommonOutRsp(errors.ErrorOK()))
}

type MemberTelAtk struct {
	Tel  string `valid:"Mobile" json:"tel"`
	Code string `valid:"Required" json:"code"`
}

// app 手机注册
type ReqRegisterMemberByTel struct {
	Tel   string `json:"tel"`
	Group int64  `json:"group"`
}

// curl -i -H "Content-Type: application/json" -i -d '{"tel":"18610889275", "code":""}' 123.57.132.7:8082/v1/hongid/register/smscode/atk
// curl -i -H "Content-Type: application/json" -i -d '{"tel":"18610889275", "code":"693512"}' 127.0.0.1:8082/v1/hongid/register/smscode/atk
// @router /register/smscode/atk [post]
func (c *RegisterController) SmsCodeAtk() {
	// 解析httpbody
	memberTelAtk := new(MemberTelAtk)
	if err := json.Unmarshal(c.getHttpBody(), memberTelAtk); err != nil {
		c.appenWrongParams(errors.NewFieldError("memberTelAtk json", err.Error()))
	}

	// 验证输入参数
	c.validation(memberTelAtk)

	// handle http request param
	if c.handleParamError() {
		return
	}

	// 验证手机验证码正确性
//	if !c.varifySmsAuthCode(memberTelAtk.Tel, memberTelAtk.Code) {
//		c.renderJson(errors.NewCommonOutRsp(errors.New(errors.CODE_BISS_ERR_SMS_CODE)))
//		return
//	}

	// 转发http请求
	reqMsg := &ReqRegisterMemberByTel{
		Tel: memberTelAtk.Tel,
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

	switch rsp.StatusCode {
	case http.StatusBadRequest, http.StatusInternalServerError:
		c.renderInternalError()
		return
	case http.StatusCreated:
		c.renderJson(errors.NewCommonOutRsp(errors.ErrorOK()))
		return
	case http.StatusOK:
		commonRsp := errors.UnmarshalCommonResponse(c.getForwardHttpBody(rsp.Body))
		if commonRsp.Code == errors.CODE_BISS_ERR_TEL_ALREADY_IN {
			c.renderJson(errors.NewCommonOutRsp(errors.New(errors.CODE_BISS_ERR_TEL_ALREADY_IN)))
			return
		}
	}

	c.renderJson(errors.NewCommonOutRsp(errors.New(errors.CODE_SYS_ERR_BASE)))
}

type MemberRegister struct {
	Tel      string `json:"tel" valid:"Mobile"`
	NickName string `json:"nickname" valid:"Required"`
	Password string `json:"password" valid:"Required"`
}

// curl -i -H "Content-Type: application/json" -i -d '{"tel":"18610889275", "nickName":"mint","password":"123456"}' 123.57.132.7:8082/v1/hongid/register
// curl -i -H "Content-Type: application/json" -i -d '{"tel":"18610889275", "nickName":"mint","password":"123456"}' 127.0.0.1:8082/v1/hongid/register
// @router /register [post]
func (c *RegisterController) Register() {
	// 解析httpbody
	memberReg := new(MemberRegister)
	if err := json.Unmarshal(c.getHttpBody(), memberReg); err != nil {
		c.appenWrongParams(errors.NewFieldError("memberTelAtk json", err.Error()))
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
	switch rspMemberInfo.StatusCode {
	case http.StatusInternalServerError:
		c.renderInternalError()
		return
	case http.StatusOK:
		break
	}

	member := new(hm.Member)
	if err := json.Unmarshal(c.getForwardHttpBody(rspMemberInfo.Body), member); err != nil {
		c.renderInternalError()
		return
	}

	member.NickName = memberReg.NickName
	member.PassWord = security.GenerateFromPassword(memberReg.Password)
	member.Status = hm.EMemberStatus_Use

	reqBytes, err := json.Marshal(member)
	if err != nil {
		c.renderInternalError()
		return
	}
	rspMemberUpd, err := c.forwardHttp("PUT", fmt.Sprintf(hongIdInfoById, hongIdHost, member.Id), bytes.NewReader(reqBytes))
	if err != nil {
		c.renderInternalError()
		return
	}
	switch rspMemberUpd.StatusCode {
	case http.StatusOK:
		c.renderJson(errors.NewCommonOutRsp(errors.ErrorOK()))
	default:
		c.renderInternalError()
	}
}
