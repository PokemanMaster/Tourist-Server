package service

import (
	"qiqi-go/models"
	"qiqi-go/response"
)

type ShowOrderService struct{}

// Show 展示订单商品内容
func (service *ShowOrderService) Show(num string) response.Response {
	var order models.Orders
	var product models.Products

	//根据id查找order
	err := models.DB.Where("order_num=?", num).First(&order).Error
	if err != nil {
		return response.Response{Status: 401, Data: nil, Msg: "", Error: err.Error()}
	}

	//根据order查找product
	err = models.DB.Where("id=?", order.ProductID).First(&product).Error
	if err != nil {
		return response.Response{Status: 402, Data: nil, Msg: "", Error: err.Error()}
	}
	return response.Response{Status: 401, Data: product, Msg: "", Error: ""}
}
