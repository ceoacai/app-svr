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
	"github.com/astaxie/beego/validation"
	"github.com/globalways/utils_go/consts"
	"github.com/globalways/utils_go/errors"
	"log"
	"regexp"
)

type PublicController struct {
	BaseController
}

type MemberTel struct {
	Tel  string `form:"tel"`
	Type int    `form:"type"`
}

// valid
func (p *MemberTel) Valid(v *validation.Validation) {
	if b, err := regexp.MatchString(consts.Regexp_Mobile, p.Tel); !b || err != nil {
		v.SetError("tel", "暂不支持此手机号码.")
	}
}

// curl -i -d "tel=18610889277" 123.57.132.7:8082/v1/public/smscode
// curl -i -d "tel=18610889277" 127.0.0.1:8082/v1/public/smscode
// @router /smscode [post]
func (c *PublicController) SmsCode() {
	// 解析httpbody
	memberTel := new(MemberTel)
	if err := c.ParseForm(memberTel); err != nil {
		c.renderInternalError()
		return
	}

	log.Printf("req: %+v\n", memberTel)

	// 验证手机号正确性
	c.validation(memberTel)

	// handle http request param
	if c.handleParamError() {
		log.Println("参数错误")
		return
	}

	// TODO 请求短信网关
	_, err := c.sendSMScode(memberTel.Tel, memberTel.Type)
	if err != nil {
		log.Println("短信发送错误")
		c.renderJson(errors.NewClientRsp(errors.CODE_BISS_ERR_SMS_GATE_FAIL))
		return
	}

	c.renderJson(errors.NewGlobalwaysErrorRsp(errors.ErrorOK()))
}
