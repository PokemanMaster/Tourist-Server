package models

import (
	"gorm.io/gorm"
)

// 用户状态
const (
	// PassWordCost 密码加密难度
	PassWordCost = 12
	// Active 激活用户
	Active string = "active"
	// Inactive 未激活用户
	Inactive string = "inactive"
	// Suspend 被封禁用户
	Suspend string = "suspend"
)

// Users 用户表
type Users struct {
	gorm.Model
	UserID    uint   `json:"userId"`    // 用户标识
	Name      string `json:"name"`      // 用户名字
	Nickname  string `json:"nickname"`  // 用户昵称
	Telephone string `json:"telephone"` // 用户电话
	Password  string `json:"password"`  // 用户密码
	Status    string `json:"status"`    // 用户状态
	Limit     int    `json:"limit"`     // 用户与某种关系的限制
	Avatar    string `json:"avatar"`    // 用户头像
}

// user是表的名字
var user Users

// GetUserInfoByAccountPassword 通过账号和密码获取用户信息
func GetUserInfoByAccountPassword(Account, Password string) (*Users, error) {
	if err := DB.Where("Account = ? AND Password = ?", Account, Password).First(&user).Error; err != nil {
		return nil, err // 返回错误
	}
	return &user, nil
}

// GetUserInfoByAccount 通过账号获取用户信息
func GetUserInfoByAccount(Account string) (*Users, error) {
	if err := DB.Where("Account = ?", Account).First(&user).Error; err != nil {
		return nil, err // 返回错误
	}
	return &user, nil
}

// GetUserInfoByPassword 通过密码获取用户信息
func GetUserInfoByPassword(Password string) (*Users, error) {
	if err := DB.Where("Password = ?", Password).First(&user).Error; err != nil {
		return nil, err // 返回错误
	}
	return &user, nil
}

// GetUserInfoByNickname 通过名字获取用户信息
func GetUserInfoByNickname(Nickname string) (*Users, error) {
	if err := DB.Where("Nickname = ?", Nickname).First(&user).Error; err != nil {
		return nil, err // 返回错误
	}
	return &user, nil
}

// IsTelephoneExists 判断用户电话号码是否存在
func IsTelephoneExists(telephone string) bool {
	DB.Where("telephone = ?", telephone).First(&user)
	return user.ID != 0
}

// IsAccountExists  判断用户账号是否存在
func IsAccountExists(Account string) bool {
	DB.Where("Account = ?", Account).First(&user)
	return user.ID != 0
}
