package service

import (
	"log"
	"qiqi-go/models"
	"qiqi-go/response"
)

// ListProductHomeService 前端传过来的商品数据
type ListProductHomeService struct {
	Name       string `json:"name"`        // 商品名字
	Label      string `json:"label"`       // 商品标签
	Start      int    `json:"start"`       // 商品分类页的起始
	CategoryID uint   `json:"category_id"` // 分类商品的id，例如：家电1
}

// HomeList 三个商品首页列表
func (service *ListProductHomeService) HomeList() response.Response {
	// 定义数据临时存储
	var ProductsHome []models.Products
	err := models.DB.Table("Products").
		Limit(3).
		Scan(&ProductsHome).Error
	if err != nil {
		log.Print(err)
		return response.Response{Status: 501, Data: nil, Msg: "", Error: err.Error()}
	}
	return response.Response{Status: 201, Data: ProductsHome, Msg: "", Error: ""}
}
