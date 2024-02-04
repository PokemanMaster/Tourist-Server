package service

import (
	"log"
	"qiqi-go/models"
	"qiqi-go/response"
)

// ListProductSortService 前端传过来的商品数据
type ListProductSortService struct {
	Label      string `json:"label"`       // 商品标签
	Start      int    `json:"start"`       // 商品分类页的起始
	CategoryID uint   `json:"category_id"` // 分类商品的id，例如：家电1
}

// SortList 商品分类列表
func (service *ListProductSortService) SortList() response.Response {
	// 定义数据临时存储
	var ProductsSort []models.Products
	err := models.DB.Table("Products").
		Where("products_label = ?", service.Label).
		Scan(&ProductsSort).Error
	if err != nil {
		log.Print(err)
		return response.Response{Status: 501, Data: nil, Msg: "", Error: err.Error()}
	}
	return response.Response{Status: 201, Data: ProductsSort, Msg: "", Error: ""}
}
