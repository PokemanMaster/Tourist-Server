package service

import (
	"qiqi-go/models"
	"qiqi-go/response"
)

// UpdateAddressService 收货地址修改的服务
type UpdateAddressService struct {
	ID      uint   `json:"id"`
	UserID  uint   `json:"user_id"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

// Update 修改购物车信息
func (service *UpdateAddressService) Update() response.Response {
	address := models.Address{
		UserID:  service.UserID,
		Name:    service.Name,
		Phone:   service.Phone,
		Address: service.Address,
	}
	address.ID = service.ID
	// 修改数据库里的收货地址
	err := models.DB.Save(&address).Error
	if err != nil {
		return response.Response{Status: 401, Data: nil, Msg: "", Error: err.Error()}
	}
	// 修改完后就重新展示
	var addresses []models.Address
	err = models.DB.Where("user_id=?", service.UserID).Order("created_at desc").Find(&addresses).Error
	if err != nil {
		return response.Response{Status: 401, Data: nil, Msg: "", Error: err.Error()}
	}
	return response.Response{Status: 201, Data: addresses, Msg: "", Error: ""}
}
