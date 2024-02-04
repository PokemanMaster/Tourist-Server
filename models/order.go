package models

import "gorm.io/gorm"

// Orders 订单模型
type Orders struct {
	gorm.Model
	UserID       uint   // 用户id
	ProductID    uint   // 商品id
	Num          uint   // 商品数量
	OrderNum     uint64 // 订单编号
	AddressName  string // 收货人姓名
	AddressPhone string // 收货人电话
	Address      string // 收货人地址
	Type         uint   // 订单状态
}
