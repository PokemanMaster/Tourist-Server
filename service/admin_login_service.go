package service

import (
	"golang.org/x/crypto/bcrypt"
	"qiqi-go/models"
	"qiqi-go/response"
)

// AdminLoginService 管理员登录的服务
type AdminLoginService struct {
	UserName string `json:"user_name" binding:"required,min=5,max=30"`
	Password string `json:"password" binding:"required,min=8,max=40"`
}

// AdminLogin 管理员登录
func (service *AdminLoginService) AdminLogin() response.Response {
	var admin models.Admin

	err := models.DB.Where("user_name = ?", service.UserName).First(&admin).Error
	if err != nil {
		return response.Response{Status: 401, Data: nil, Msg: "", Error: ""}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(service.Password), []byte(service.Password)); err != nil {
		return response.Response{Status: 501, Data: nil, Msg: "您输入的密码错误", Error: ""}
	}

	return response.Response{Status: 201, Data: nil, Msg: "ok", Error: ""}
}
