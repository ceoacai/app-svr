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

type StoreController struct {
	BaseController
}

type ReqStoreList struct {
	GPS string `form:"gps"`

	OrderType    int  `form:"ordertype"`
	ProductCount int  `form:"productcount"`
	StorePage    uint `form:"storepage"`
	StoreSize    uint `form:"storesize"`

	IndustryId int64 `form:"industryid"`
	Distance   int   `form:"distance"`

	KeywordSearch string `form:"keywordsearch"`
}

type BodyStoreBrushFilters struct {
	IndustryFilters []*BrushFilterItem `json:"industryfilters"`
	DistanceFilters []*BrushFilterItem `json:"distancefilters"`
}

type BrushFilterItem struct {
	ItemId   int64  `json:"itemid"`
	ItemName string `json:"itemname"`
}

type BodyStoreList struct {
	Stores  []*BodyStoreBref       `json:"stores"`
	Filters *BodyStoreBrushFilters `json:"filters"` // page = 1 有值
}

type BodyStoreBref struct {
	StoreId      int64              `json:"storeid"`
	StoreName    string             `json:"storename"`
	IndustryName string             `json:"industryname"`
	Distance     int                `json:"distance"`
	Avatar       string             `json:"avatar"`
	Products     []*BodyProductBref `json:"products"`
}

type BodyProductBref struct {
	ProductId     int64  `json:"productid"`
	ProductName   string `json:"productname"`
	ProductAvatar string `json:"productavatar"`
	ProductPrice  uint   `json:"productprice"`
	ProductUnit   string `json:"productunit"`
}

// 筛选
// @router /brush [post]
func (c *StoreController) Brush() {
}
