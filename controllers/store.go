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
	"app-svr/models"
	"github.com/globalways/utils_go/errors"
)

type StoreController struct {
	BaseController
}

type BodyStoreBrushFilters struct {
	IndustryFilters []*BrushFilterItem `json:"industryfilters"`
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
	Distance     float64            `json:"distance"`
	Avatar       string             `json:"avatar"`
	Products     []*BodyProductBref `json:"products"`
}

type BodyProductBref struct {
	ProductId       int64   `json:"productid"`
	ProductName     string  `json:"productname"`
	ProductAvatar   string  `json:"productavatar"`
	ProductPrice    float64 `json:"productprice"`
	ProductCurrency string  `json:"productcurrency"`
	ProductUnit     string  `json:"productunit"`
}

// 筛选
// curl -i -d "gps=4938,8473&ordertype=1&orderorder=1&productcount=3&storepage=1&storesize=10&industryid=0&distance=1000" 127.0.0.1:8082/v1/store/brush
// @router /brush [post]
func (c *StoreController) BrushStoreList() {
	reqMsg := new(models.ReqStoreList)
	if err := c.ParseForm(reqMsg); err != nil {
		c.renderInternalError()
		return
	}

	// 参数验证
	c.validation(reqMsg)

	// handle param error
	if c.handleParamError() {
		return
	}

	stores := models.SearchStore(reqMsg)
	storesBref := make([]*BodyStoreBref, 0)
	for _, store := range stores {
		storeBref := &BodyStoreBref{
			StoreId:      store.StoreId,
			StoreName:    store.StoreName,
			IndustryName: store.IndustryName,
			Avatar:       store.Avatar,
			Products: func() []*BodyProductBref {
				products := make([]*BodyProductBref, 0)
				for _, product := range store.Products {
					productBref := &BodyProductBref{
						ProductId:     product.ProductId,
						ProductName:   product.ProductName,
						ProductAvatar: product.ProductAvatar,
						ProductPrice:  product.ProductPrice,
						ProductUnit:   product.ProductUnit,
					}
					products = append(products, productBref)
				}
				return products
			}(),
		}
		storesBref = append(storesBref, storeBref)
	}

	clientRsp := new(errors.ClientRsp)
	clientRsp.Status = errors.NewStatus(errors.CODE_SUCCESS)
	clientRsp.Body = &BodyStoreList{
		Stores: storesBref,
		Filters: func() *BodyStoreBrushFilters {
			if reqMsg.StorePage != 1 {
				return nil
			}

			filters := new(BodyStoreBrushFilters)

			industryFilters := make([]*BrushFilterItem, 0)
			for key, val := range models.Industrys {
				industryFilters = append(industryFilters, &BrushFilterItem{key, val})
			}
			filters.IndustryFilters = industryFilters

			return filters
		}(),
	}

	c.renderJson(clientRsp)
}

type ReqStoreDetail struct {
	StoreId int64 `form:"storeid"`
}

type BodyStoreDetail struct {
	StoreId         int64  `json:"storeid"`
	StoreName       string `json:"storename"`
	StoreDesc       string `json:"storedesc"`
	StoreGPS        string `json:"storegps"`
	StoreAddress    string `json:"storeaddress"`
	Avatar          string `json:"avatar"`
	StorePhone      string `json:"storephone"`
	ProductHotLimit uint   `json:"producthotlimit"`
}

// @router /storeDetail [post]
func (c *StoreController) StoreDetail() {

}

type ReqStoreProducts struct {
	StoreId     int64 `form:"storeid"`
	ProductPage uint  `form:"productpage"`
	ProductSize uint  `form:"productsize"`
}

type BodyStoreProducts struct {
	ProductId         int64   `json:"productid"`
	ProductName       string  `json:"productname"`
	ProductAvatar     string  `json:"productavatar"`
	ProductPrice      float64 `json:"productprice"`
	ProductCurrency   string  `json:"productcurrency"`
	ProductUnit       string  `json:"productunit"`
	ProductSalesCnt   uint    `json:"productsalescnt"`
	ProductStockCnt   uint    `json:"productstockcnt"`
	ProductViewCnt    uint    `json:"productviewcnt"`
	MgrRecommendFlag  bool    `json:"mgrrecommendflag"`
	ProductStockLimit uint    `json:"productstocklimit"`
}

// @router /products [post]
func (c *StoreController) StoreProductList() {

}
