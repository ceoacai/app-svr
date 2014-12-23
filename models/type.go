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
package models


// 验证码类型
const (
	TYPE_SMS_CODE_REGISTER = iota + 1
	TYPE_SMS_CODE_FIND_PASSWORD
)

// 商铺排序类型
const (
	TYPE_STORE_ORDER_RECOMMEND = iota + 1 //综合
	TYPE_STORE_ORDER_SALES                // 所有商品销量
	TYPE_STORE_ORDER_HOTS                 // 逛的会员最多
)

// 高低排序
const (
	TYPE_ORDER_ASC = iota + 1
	TYPE_ORDER_DESC
)

// 商品排序类型
const (
	TYPE_PRODUCT_ORDER_RECOMMEND = iota + 1 //综合
	TYPE_PRODUCT_ORDER_SALES                // 商品销量
	TYPE_PRODUCT_ORDER_PRICE                // 商品价格
)

