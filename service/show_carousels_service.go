package service

import (
	"log"
	"qiqi-go/models"
	"qiqi-go/response"
)

// ShowCarouselsService 前端传过来的商品数据
type ShowCarouselsService struct{}

// GetCarouselsImages 获取轮播图
func (service *ShowCarouselsService) GetCarouselsImages() response.Response {
	// 定义数据临时存储
	var carousels []models.Carousels
	err := models.DB.Find(&carousels).Error
	if err != nil {
		log.Print(err)
		return response.Response{Status: 501, Data: nil, Msg: "", Error: err.Error()}
	}
	return response.Response{Status: 201, Data: carousels, Msg: "", Error: ""}
}
