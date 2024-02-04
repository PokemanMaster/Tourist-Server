package service

import (
	"qiqi-go/models"
	"qiqi-go/response"
)

// CreateAddressService 收货地址创建的服务
type CreateAddressService struct {
	UserID  uint   `json:"user_id"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

// Create 用户创建收货地址，同时展示自己已经创建过的地址
func (service *CreateAddressService) Create() response.Response {
	// 接受前端传入数据保存到address
	address := models.Address{
		UserID:  service.UserID,
		Name:    service.Name,
		Phone:   service.Phone,
		Address: service.Address,
	}
	// 创建收货地址
	err := models.DB.Create(&address).Error
	if err != nil {
		return response.Response{Status: 401, Data: nil, Msg: "", Error: err.Error()}
	}
	// 同时展示自己已经创建的收货地址
	var addresses []models.Address
	err = models.DB.Where("user_id=?", service.UserID).Order("created_at desc").Find(&addresses).Error
	if err != nil {
		return response.Response{Status: 401, Data: nil, Msg: "", Error: err.Error()}
	}
	return response.Response{Status: 201, Data: addresses, Msg: "", Error: ""}
}
