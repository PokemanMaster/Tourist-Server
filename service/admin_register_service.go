package service

import (
	"golang.org/x/crypto/bcrypt"
	"qiqi-go/models"
	"qiqi-go/response"
)

// AdminRegisterService 管理用户注册服务
type AdminRegisterService struct {
	UserName        string `json:"user_name" binding:"required,min=5,max=30"`
	Password        string `json:"password" binding:"required,min=8,max=40"`
	PasswordConfirm string `json:"password_confirm" binding:"required,min=8,max=40"`
}

// AdminRegister 管理员注册
func (services *AdminRegisterService) AdminRegister() response.Response {
	if services.UserName == "" {
		return response.Response{Status: 401, Data: nil, Msg: "", Error: ""}
	}

	// 判断用户是否存在
	var count int64
	models.DB.Model(&models.Admin{}).Where("user_name = ?", services.UserName).Count(&count)
	if count > 0 {
		return response.Response{Status: 401, Data: nil, Msg: "", Error: ""}
	}

	// 密码重新确认
	if services.PasswordConfirm != services.Password {
		return response.Response{Status: 401, Data: nil, Msg: "", Error: ""}
	}

	hasePassword, err := bcrypt.GenerateFromPassword([]byte(services.Password), bcrypt.DefaultCost)
	if err != nil {
		return response.Response{Status: 501, Data: nil, Msg: "加密失败", Error: ""}
	}

	newAdmin := models.Admin{
		UserName:       services.UserName,
		PasswordDigest: string(hasePassword),
		Avatar:         models.Active,
	}

	err = models.DB.Create(&newAdmin).Error
	if err != nil {
		return response.Response{Status: 501, Data: nil, Msg: "加密失败", Error: ""}
	}

	return response.Response{Status: 201, Data: nil, Msg: "ok", Error: ""}
}
