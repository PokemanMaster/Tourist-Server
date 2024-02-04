package service

import (
	"log"
	"qiqi-go/models"
	"qiqi-go/response"
)

// ListProductSearchService 前端传过来的商品数据
type ListProductSearchService struct {
	Name string `json:"name"` // 商品名字
}

// SortListSearch 商品分类页搜索
func (service *ListProductSearchService) SortListSearch() response.Response {
	// 定义数据临时存储
	var ProductsSort []models.Products
	err := models.DB.Table("Products").
		Where("products_name LIKE ?", "%"+service.Name+"%").
		Scan(&ProductsSort).Error
	if err != nil {
		log.Print(err)
		return response.Response{Status: 501, Data: nil, Msg: "", Error: err.Error()}
	}
	return response.Response{Status: 201, Data: ProductsSort, Msg: "", Error: ""}
}
