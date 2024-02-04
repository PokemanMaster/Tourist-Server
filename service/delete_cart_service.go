package service

import (
	"fmt"
	"qiqi-go/cache"
	"qiqi-go/response"
)

// DeleteCartService 前端传递过来的数据
type DeleteCartService struct {
	ProductId int  // 商品id
	UserId    uint // 用户id
}

// Delete 移除购物车的商品
func (service *DeleteCartService) Delete() response.Response {

	if service == nil {
		return response.Response{Status: 500, Data: nil, Msg: "DeleteCartService is nil", Error: ""}
	}

	cartKey := fmt.Sprintf("%s%d", "cart:", service.UserId)
	if cache.RDB == nil {
		return response.Response{Status: 500, Data: nil, Msg: "Redis client is nil", Error: ""}
	}

	err := cache.RDB.HDel(cartKey, fmt.Sprintf("%d", service.ProductId)).Err()
	if err != nil {
		return response.Response{Status: 501, Data: nil, Msg: "删除数据失败", Error: err.Error()}
	}

	return response.Response{Status: 201, Data: nil, Msg: "ok", Error: ""}

	//var cart []models.Cart
	//err := models.DB.Table("cart").
	//	Where("user_id = ? AND product_id = ?", service.UserID, service.ProductID).Find(&cart).Error
	//if err != nil {
	//	return response.Response{Status: 501, Data: nil, Msg: "删除数据失败", Error: err.Error()}
	//}
	//
	//err = models.DB.Delete(&cart).Error
	//if err != nil {
	//	return response.Response{Status: 501, Data: nil, Msg: "删除数据失败", Error: err.Error()}
	//}
}
