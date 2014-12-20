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
	"github.com/globalways/utils_go/errors"
	"fmt"
	"net/http"
)

type MemberCardController struct {
	BaseController
}

// curl -i 127.0.0.1:8082/v1/memberCards/card/6320860000000000011/bind/1
// @router /card/:card/bind/:owner [post]
func (c *MemberCardController) BindCard() {
	cardStr := c.GetString(":card")
	ownerId, err1 := c.GetInt64(":owner")
	if err1 != nil {
		c.appenWrongParams(errors.NewFieldError(":owner", err1.Error()))
	}

	// handle http error
	if c.handleParamError() {return}

	rsp, err := c.forwardHttp("POST", fmt.Sprintf(memberCardBind, hongIdHost, cardStr, ownerId), nil)
	if err != nil || rsp.StatusCode != http.StatusOK {
		c.renderInternalError()
		return
	}

	commonRsp := errors.UnmarshalCommonResponse(c.getForwardHttpBody(rsp.Body))
	if commonRsp.Code == errors.CODE_DB_ERR_UPDATE {
		c.renderInternalError()
		return
	}

	c.renderJson(commonRsp)
}

// curl -i 127.0.0.1:8082/v1/memberCards/card/6320860000000000011/unbind
// @router /card/:card/unbind [post]
func (c *MemberCardController) UnBindCard() {
	cardStr := c.GetString(":card")

	// handle http error
	if c.handleParamError() {return}

	rsp, err := c.forwardHttp("POST", fmt.Sprintf(memberCardUnBind, hongIdHost, cardStr), nil)
	if err != nil || rsp.StatusCode != http.StatusOK {
		c.renderInternalError()
		return
	}

	commonRsp := errors.UnmarshalCommonResponse(c.getForwardHttpBody(rsp.Body))
	if commonRsp.Code == errors.CODE_DB_ERR_UPDATE {
		c.renderInternalError()
		return
	}

	c.renderJson(commonRsp)
}
