package service

import (
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"qiqi-go/cache"
	"qiqi-go/response"
)

// CreateCartService 前端传递过来的数据
type CreateCartService struct {
	UserId              uint            // 用户id
	ProductId           int             // 商品id
	ProductsImage       string          // 商品图片
	ProductsName        string          // 商品名字
	ProductsLabel       string          // 商品标签
	ProductsTitle       string          // 商品标题，用于展示一级详情页
	ProductsDescription string          // 商品描述，用于展示二级详情页
	ProductsMoney       decimal.Decimal // 商品金额
	Num                 uint            // 商品数量
}

// Create 添加商品到购物车
func (service *CreateCartService) Create() response.Response {
	//用用户id作为key
	cartKey := fmt.Sprintf("%s%d", "cart:", service.UserId)

	productData := struct {
		ProductId           int
		ProductsImage       string
		ProductsName        string
		ProductsLabel       string
		ProductsTitle       string
		ProductsDescription string
		ProductsMoney       string
		Num                 uint
	}{
		service.ProductId,
		service.ProductsImage,
		service.ProductsName,
		service.ProductsLabel,
		service.ProductsTitle,
		service.ProductsDescription,
		service.ProductsMoney.String(), // Convert decimal to string
		service.Num,
	}

	// 将结构体转换为 JSON 字符串
	productJSON, err := json.Marshal(productData)
	if err != nil {
		return response.Response{Status: 500, Data: nil, Msg: "Internal Server Error", Error: err.Error()}
	}

	// 使用 HSet 将 JSON 字符串存储在 Redis 中
	err = cache.RDB.HSet(cartKey, fmt.Sprintf("%d", service.ProductId), productJSON).Err()
	if err != nil {
		return response.Response{Status: 401, Data: nil, Msg: "创建购物车失败", Error: err.Error()}
	}
	return response.Response{Status: 201, Data: nil, Msg: "ok", Error: ""}

	//cart, err := json.Marshal(models.Cart{
	//	UserID:    service.UserID,
	//	ProductID: service.ProductID,
	//	Num:       service.Num,
	//})
	//err = models.DB.Create(&cart).Error
	//if err != nil {
	//	return response.Response{Status: 500, Data: nil, Msg: "Internal Server Error", Error: err.Error()}
	//}
}
