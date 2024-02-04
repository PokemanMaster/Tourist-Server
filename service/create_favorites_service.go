package service

import (
	"qiqi-go/models"
	"qiqi-go/response"
)

// 添加商品到收藏夹

// CreateFavoriteService 收藏创建的服务
type CreateFavoriteService struct {
	UserID    uint `json:"user_id"`
	ProductID uint `json:"product_id"`
}

// Create 创建收藏夹
func (service *CreateFavoriteService) Create() response.Response {
	var favorite models.Favorites

	models.DB.Where("user_id=? AND product_id=?", service.UserID, service.ProductID).Find(&favorite)
	if favorite == (models.Favorites{}) {
		favorite = models.Favorites{
			UserID:    service.UserID,
			ProductID: service.ProductID,
		}
		if err := models.DB.Create(&favorite).Error; err != nil {
			return response.Response{Status: 401, Data: nil, Msg: "", Error: err.Error()}
		}
	} else {
		return response.Response{Status: 402, Data: nil, Msg: "", Error: ""}
	}
	return response.Response{Status: 201, Data: nil, Msg: "", Error: ""}
}
