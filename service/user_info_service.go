package service

import (
	"qiqi-go/middleware"
	"qiqi-go/models"
	"qiqi-go/response"
)

// UserInfoService 前端请求过来的数据
type UserInfoService struct {
	Token string `json:"token"` // 用户token
}

// UserInfo 获取用户信息
func (service *UserInfoService) UserInfo() response.Response {
	// 从token中获取电话
	_, b, _ := middleware.ParseToken(service.Token)
	telephone := b.Telephone
	var UserInfo []models.Users
	err := models.DB.Table("users").Where("users.telephone = ?", telephone).Scan(&UserInfo).Error
	if err != nil {
		return response.Response{Status: 501, Data: nil, Msg: "系统异常", Error: ""}
	}
	return response.Response{Status: 201, Data: UserInfo, Msg: "", Error: ""}
}
