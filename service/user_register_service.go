package service

import (
	"golang.org/x/crypto/bcrypt"
	"qiqi-go/middleware"
	"qiqi-go/middleware/midAuto"
	"qiqi-go/middleware/midUtil"
	"qiqi-go/models"
	"qiqi-go/response"
)

// UserRegisterService 前端请求过来的数据
type UserRegisterService struct {
	Name      string `json:"name"`      // 用户名字
	Telephone string `json:"telephone"` // 用户电话
	Password  string `json:"password"`  // 用户密码
	Code      string `json:"code"`      // 用户输入的验证码
	CodeId    string `json:"codeId"`    // 正确的验证码
}

// UserRegister 用户注册
func (service *UserRegisterService) UserRegister(userID, status interface{}) response.Response {
	//获取参数
	name := service.Name
	telephone := service.Telephone
	password := service.Password
	code := service.Code
	codeId := service.CodeId // 验证验证码正确的id
	//数据验证
	if len(name) == 0 {
		return response.Response{Status: 501, Data: nil, Msg: "名字不能为空", Error: ""}
	}
	// 判断手机是否正确
	if !midUtil.TelephoneNumberIsTure(telephone) {
		return response.Response{Status: 501, Data: nil, Msg: "手机号必须为11位且符合+86", Error: ""}
	}
	// 判断密码是否正确
	if !midUtil.PasswordIsTure(password) {
		return response.Response{Status: 501, Data: nil, Msg: "密码不能少于6位且由数字和字母组合", Error: ""}
	}
	//判断手机号码是否存在
	if models.IsTelephoneExists(telephone) {
		return response.Response{Status: 501, Data: nil, Msg: "用户已经存在", Error: ""}
	}
	// 校验验证码答案
	answer := midAuto.GetCodeAnswer(codeId)
	if code != answer {
		return response.Response{Status: 501, Data: nil, Msg: "你输入的验证码错误", Error: ""}
	}
	//给密码做hash算法加密处理
	hasePassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return response.Response{Status: 501, Data: nil, Msg: "加密失败", Error: ""}
	}
	// 生成userId
	userId := midAuto.GenerateRandomNumber()
	if err != nil {
		return response.Response{Status: 501, Data: nil, Msg: "userId生成失敗", Error: ""}
	}
	//把用户信息放入数据表
	newUser := models.Users{
		UserID:    userId,
		Name:      name,
		Telephone: telephone,
		Password:  string(hasePassword),
		Status:    models.Active,
	}
	err = models.DB.Table("users").Create(&newUser).Error
	if err != nil {
		return response.Response{Status: 501, Data: nil, Msg: "系统错误，创建用户失败", Error: ""}
	}
	//给用户发放token
	token, err := middleware.ReleaseToken(newUser)
	if err != nil {
		return response.Response{Status: 501, Data: nil, Msg: "系统异常", Error: ""}
	}
	return response.Response{Status: 201, Data: token, Msg: "", Error: ""}
}
