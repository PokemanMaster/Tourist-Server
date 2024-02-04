package service

import (
	"qiqi-go/models"
	"qiqi-go/response"
)

// ShowFavoritesService 展示收藏夹详情的服务
type ShowFavoritesService struct {
	Limit int `form:"limit"`
	Start int `form:"start"`
}

// Show 商品图片
func (service *ShowFavoritesService) Show(id string) response.Response {
	var favorites []models.Favorites

	err := models.DB.Model(&favorites).Where("user_id=?", id).Error
	if err != nil {
		return response.Response{Status: 401, Data: nil, Msg: "", Error: err.Error()}
	}

	err = models.DB.Where("user_id=?", id).Limit(service.Limit).Offset(service.Start).Find(&favorites).Error
	if err != nil {
		return response.Response{Status: 401, Data: nil, Msg: "", Error: err.Error()}
	}
	return response.Response{Status: 401, Data: favorites, Msg: "", Error: ""}
}
