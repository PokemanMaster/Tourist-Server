package service

import (
	"qiqi-go/middleware"
	"qiqi-go/models"
	"qiqi-go/response"
)

// UserEditService 前端请求过来的数据
type UserEditService struct {
	NickName string `json:"nickName"` // 用户别名
	Name     string `json:"name"`     // 用户名字
	Password string `json:"password"` // 用户密码
	Avatar   string `json:"avatar"`   // 用户头像
	Token    string `json:"token"`    // 用户token
}

// UserEdit 用户更新修改信息
func (service *UserEditService) UserEdit() response.Response {
	_, b, _ := middleware.ParseToken(service.Token)
	telephone := b.Telephone
	updateData := UserEditService{
		NickName: service.NickName,
		Name:     service.Name,
		Password: service.Password,
		Avatar:   service.Avatar,
	}
	err := models.DB.Table("users").Where("users.telephone = ?", telephone).Updates(updateData).Error
	if err != nil {
		return response.Response{Status: 501, Data: nil, Msg: "系统异常", Error: ""}
	}
	return response.Response{Status: 201, Data: nil, Msg: "ok", Error: ""}
}
