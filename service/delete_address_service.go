package service

import (
	"qiqi-go/models"
	"qiqi-go/response"
)

// DeleteAddressService 购物车删除的服务
type DeleteAddressService struct {
	AddressID uint `json:"address_id"`
}

// Delete 删除收货地址
func (service *DeleteAddressService) Delete() response.Response {
	var address models.Address

	err := models.DB.Where("id = ?", service.AddressID).Find(&address).Error
	if err != nil {
		return response.Response{Status: 401, Data: nil, Msg: "", Error: err.Error()}
	}

	err = models.DB.Delete(&address).Error
	if err != nil {
		return response.Response{Status: 401, Data: nil, Msg: "", Error: err.Error()}
	}
	return response.Response{Status: 201, Data: nil, Msg: "", Error: ""}
}
