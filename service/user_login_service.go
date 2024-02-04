package service

import (
	"golang.org/x/crypto/bcrypt"
	"qiqi-go/middleware"
	"qiqi-go/middleware/midAuto"
	"qiqi-go/models"
	"qiqi-go/response"
	"qiqi-go/response/BuildUser"
)

// UserLoginService 前端请求过来的数据
type UserLoginService struct {
	Telephone string `json:"telephone"` // 用户电话
	Password  string `json:"password"`  // 用户密码
	Code      string `json:"code"`      // 用户输入的验证码
	CodeId    string `json:"codeId"`    // 正确的验证码
}

// UserLogin 用户登录
func (service *UserLoginService) UserLogin(userID, status interface{}) response.Response {
	//判断手机号是否存在, 查询的是user表
	var user models.Users
	err := models.DB.Where("telephone = ?", service.Telephone).First(&user).Error
	if err != nil {
		return response.Response{Status: 501, Data: nil, Msg: "电话或密码不能为空", Error: ""}
	}
	if user.ID == 0 && user.Telephone == "" && user.Password == "" {
		return response.Response{Status: 501, Data: nil, Msg: "没有查询到此手机号", Error: ""}
	}
	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(service.Password)); err != nil {
		return response.Response{Status: 501, Data: nil, Msg: "您输入的密码错误", Error: ""}
	}
	// 检测验证码是否正确
	answer := midAuto.GetCodeAnswer(service.CodeId)
	if service.Code != answer {
		return response.Response{Status: 501, Data: nil, Msg: "你输入的验证码错误", Error: ""}
	}
	//发放token
	token, err := middleware.ReleaseToken(user)
	if err != nil {
		return response.Response{Status: 501, Data: nil, Msg: "系统异常", Error: ""}
	}
	return response.Response{Status: 201, Data: BuildUser.BuildUser{Token: token, UserId: user.UserID}, Msg: "登录成功", Error: ""}
}
