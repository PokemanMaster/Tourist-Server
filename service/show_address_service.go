package service

import (
	"qiqi-go/models"
	"qiqi-go/response"
)

// ShowAddressService 展示收货地址的服务
type ShowAddressService struct{}

// Show 展示用户的收货地址
func (service *ShowAddressService) Show(id string) response.Response {
	var addresses []models.Address
	err := models.DB.Where("user_id=?", id).Order("created_at desc").Find(&addresses).Error
	if err != nil {
		return response.Response{Status: 401, Data: nil, Msg: "", Error: err.Error()}
	}
	return response.Response{Status: 201, Data: addresses, Msg: "", Error: ""}
}
