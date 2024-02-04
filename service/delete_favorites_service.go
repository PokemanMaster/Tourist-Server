package service

import (
	"qiqi-go/models"
	"qiqi-go/response"
)

// DeleteFavoritesService 删除收藏的服务
type DeleteFavoritesService struct {
	UserID    uint `form:"user_id" json:"user_id"`
	ProductID uint `form:"product_id" json:"product_id"`
}

// Delete 删除收藏
func (service *DeleteFavoritesService) Delete() response.Response {
	var favorite models.Favorites

	err := models.DB.Where("user_id=? AND product_id=?", service.UserID, service.ProductID).Find(&favorite).Error
	if err != nil {
		return response.Response{Status: 401, Data: nil, Msg: "", Error: err.Error()}
	}

	err = models.DB.Delete(&favorite).Error
	if err != nil {
		return response.Response{Status: 401, Data: nil, Msg: "", Error: err.Error()}
	}

	return response.Response{Status: 201, Data: nil, Msg: "", Error: ""}
}
