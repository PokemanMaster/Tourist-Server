package service

import (
	"qiqi-go/models"
	"qiqi-go/response"
)

// 商品类

// ShowProductService 前端传过来的商品数据
type ShowProductService struct {
	Name       string `json:"name"`        // 商品名字
	Label      string `json:"label"`       // 商品标签
	Start      int    `json:"start"`       // 商品分类页的起始
	CategoryID uint   `json:"category_id"` // 分类商品的id，例如：家电1
}

// Show 展示商品详情页
func (service *ShowProductService) Show(productsId string) response.Response {
	// 定义数据临时存储
	var showProducts models.Products
	err := models.DB.First(&showProducts, productsId).Error
	if err != nil {
		return response.Response{Status: 501, Data: nil, Msg: "", Error: err.Error()}
	}
	// 增加商品点击数
	showProducts.AddView()
	return response.Response{Status: 201, Data: showProducts, Msg: "", Error: ""}
}
