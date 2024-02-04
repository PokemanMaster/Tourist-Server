package models

import "gorm.io/gorm"

// Admin 管理员模型
type Admin struct {
	gorm.Model
	UserName       string
	PasswordDigest string
	Avatar         string `gorm:"size:1000"`
}
