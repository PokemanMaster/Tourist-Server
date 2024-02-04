package models

import "gorm.io/gorm"

// Carousels 轮播图表
type Carousels struct {
	gorm.Model
	Images string `json:"images"` // 图片链接地址
}
