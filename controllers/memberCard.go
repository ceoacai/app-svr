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
	"fmt"
	"net/http"
)

type MemberCardController struct {
	BaseController
}

type ReqBindCard struct {
	Owner  int64  `json:"owner"`
	Card string `json:"card"`
}

// @router /bind [post]
func (c *MemberCardController) BindCard() {
	reqMsg := new(ReqBindCard)
	if err := json.Unmarshal(c.getHttpBody(), reqMsg); err != nil {
		c.appenWrongParams(errors.NewFieldError("reqBody", err.Error()))
	}

	// handle http error
	if c.handleParamError() {return}

	rsp, err := c.forwardHttp("POST", fmt.Sprintf(memberCardBind, hongIdHost, reqMsg.Card, reqMsg.Id), nil)
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
