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

import (
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/globalways/utils_go/random"
	"sort"
	"strings"
)

type ReqStoreList struct {
	GPS string `form:"gps"`

	OrderType    int  `form:"ordertype"`
	OrderOrder   int  `form:"orderorder"`
	ProductCount int  `form:"productcount"`
	StorePage    uint `form:"storepage"`
	StoreSize    uint `form:"storesize"`

	IndustryId int64 `form:"industryid"`
	Distance   uint  `form:"distance"` // 距离xx千米以内的商铺

	KeywordSearch string `form:"keywordsearch"`
}

func (p *ReqStoreList) Valid(v *validation.Validation) {
	if len(p.GPS) == 0 {
		v.SetError("gps", "GPS信息不能为空.")
	}
	if strings.Index(p.GPS, ",") < 0 {
		v.SetError("gps", "GPS信息格式不正确.")
	}

	if p.OrderType != TYPE_STORE_ORDER_RECOMMEND &&
		p.OrderType != TYPE_STORE_ORDER_SALES &&
		p.OrderType != TYPE_STORE_ORDER_HOTS {
		v.SetError("ordertype", "排序类型错误.")
	}

	if p.OrderOrder != TYPE_ORDER_ASC && p.OrderOrder != TYPE_ORDER_DESC {
		v.SetError("orderorder", "排序顺序错误.")
	}

	if p.ProductCount < 0 {
		v.SetError("productcount", "首页商铺商品显示个数格式错误.")
	}

	if p.StorePage < 0 {
		v.SetError("storepage", "商铺分页错误.")
	}

	if p.StoreSize <= 0 {
		v.SetError("storesize", "商铺分页数量错误.")
	}

	if p.IndustryId < 0 {
		v.SetError("industryid", "商铺行业分类错误.")
	}

	if p.Distance <= 0 {
		v.SetError("distance", "商铺筛选距离错误.")
	}

}

type DataProduct struct {
	ProductId       int64
	ProductName     string
	ProductAvatar   string
	ProductPrice    float64
	ProductCurrency string
	ProductUnit     string
}

type DataStore struct {
	StoreId      int64
	StoreName    string
	IndustryName string
	GPS          string
	Avatar       string
	Products     []*DataProduct
	SalesCnt     uint
	ClickCnt     uint
}

type DataIndustry struct {
	Id   int64
	Name string
}

var (
	Stores    map[int64]*DataStore
	Industrys map[int64]string
)

func init() {

	Industrys = make(map[int64]string)
	Industrys[0] = "全部"
	Industrys[1] = "餐饮"
	Industrys[2] = "水果"

	Stores = make(map[int64]*DataStore)

	for i := int64(1); i <= 100; i++ {
		products := make([]*DataProduct, 0)
		for j := (i-1)*10 + 1; j <= i*10; j++ {
			product := &DataProduct{
				ProductId:     j,
				ProductName:   fmt.Sprintf("辣条%v", j),
				ProductAvatar: "http://img07.huishangbao.com/file/upload/201412/21/13/13-40-15-38-273480.png",
				ProductPrice:  10,
				ProductUnit:   "袋",
			}
			products = append(products, product)
		}
		store := &DataStore{
			StoreId:      i,
			StoreName:    fmt.Sprintf("辣条馆%v", i),
			IndustryName: "餐饮",
			GPS:          fmt.Sprintf("%v.%v,%v.%v", random.RandomInt(1, 180), random.RandomInt64(100000, 999999), random.RandomInt(1, 180), random.RandomInt64(100000, 999999)),
			Avatar:       "http://img4.duitang.com/uploads/item/201404/11/20140411023114_TzkKU.jpeg",
			Products:     products,
			SalesCnt:     random.RandomUint(100, 1000),
			ClickCnt:     random.RandomUint(1000, 10000),
		}

		Stores[i] = store
	}

	for i := int64(101); i <= 200; i++ {
		products := make([]*DataProduct, 0)
		for j := (i-1)*8 + 1; j <= i*8; j++ {
			product := &DataProduct{
				ProductId:     j,
				ProductName:   fmt.Sprintf("苹果%v", j),
				ProductAvatar: "http://pic27.nipic.com/20130202/11664993_180053398194_2.jpg",
				ProductPrice:  10,
				ProductUnit:   "个",
			}
			products = append(products, product)
		}
		store := &DataStore{
			StoreId:      i,
			StoreName:    fmt.Sprintf("水果店%v", i),
			IndustryName: "水果",
			GPS:          fmt.Sprintf("%v.%v,%v.%v", random.RandomInt(1, 180), random.RandomInt64(100000, 999999), random.RandomInt(1, 180), random.RandomInt64(100000, 999999)),
			Avatar:       "http://pic2.huitu.com/res/20120208/1370_20120208044334022200_1.jpg",
			Products:     products,
			SalesCnt:     random.RandomUint(100, 1000),
			ClickCnt:     random.RandomUint(1000, 10000),
		}

		Stores[i] = store
	}
}

func SearchStore(req *ReqStoreList) (stores []*DataStore) {

	if len(req.KeywordSearch) != 0 {
		for _, val := range Stores {
			if strings.Contains(val.StoreName, req.KeywordSearch) {
				stores = append(stores, val)
			} else {
				for _, product := range val.Products {
					if strings.Contains(product.ProductName, req.KeywordSearch) {
						stores = append(stores, val)
					}
				}
			}
		}
	} else {
		for _, val := range Stores {
			stores = append(stores, val)
		}
	}

	// GPS

	// OrderType
	switch req.OrderType {
	case TYPE_STORE_ORDER_HOTS:
		if req.OrderOrder == TYPE_ORDER_ASC {
			sort.Sort(StoreClickSorter(stores))
		} else {
			sort.Reverse(StoreClickSorter(stores))
		}
	case TYPE_STORE_ORDER_SALES:
	default:
	}

	//ProductCount
	for idx, val := range stores {
		if len(val.Products) > req.ProductCount {
			stores[idx].Products = stores[idx].Products[0:req.ProductCount]
		}
	}

	// sort by id
	sort.Sort(StoreIdSorter(stores))

	// StorePage StoreSize
	stores = stores[(req.StorePage-1)*req.StoreSize : req.StorePage*req.StoreSize]

	return
}

type StoreClickSorter []*DataStore

func (s StoreClickSorter) Len() int {
	return len(s)
}
func (s StoreClickSorter) Less(i, j int) bool {
	return s[i].ClickCnt > s[j].ClickCnt
}
func (s StoreClickSorter) Swap(i, j int) {
	s[i] = s[j]
}

type StoreIdSorter []*DataStore

func (s StoreIdSorter) Len() int {
	return len(s)
}
func (s StoreIdSorter) Less(i, j int) bool {
	return s[i].StoreId > s[j].StoreId
}
func (s StoreIdSorter) Swap(i, j int) {
	s[i] = s[j]
}
