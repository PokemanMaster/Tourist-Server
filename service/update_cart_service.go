package service

import (
	"encoding/json"
	"fmt"
	"qiqi-go/cache"
	"qiqi-go/response"
)

// UpdateCartService 购物车修改的服务
type UpdateCartService struct {
	UserId    uint
	ProductId int
	Num       uint
}

// Update 修改购物车信息
func (service *UpdateCartService) Update() response.Response {

	cartKey := fmt.Sprintf("%s:%d", "cart", service.UserId)
	field := fmt.Sprintf("%d", service.ProductId)

	// Get current cart data from cache
	currentData, err := cache.RDB.HGet(cartKey, field).Result()
	if err != nil {
		// Handle error when retrieving cart data
		return response.Response{Status: 401, Data: nil, Msg: "Error retrieving cart data", Error: err.Error()}
	}

	// Unmarshal JSON string to a map
	var cartData map[string]interface{}
	err = json.Unmarshal([]byte(currentData), &cartData)
	if err != nil {
		// Handle JSON unmarshal error
		return response.Response{Status: 401, Data: nil, Msg: "Error unmarshaling JSON", Error: err.Error()}
	}

	// Update the "Num" property with the value from service.Num
	cartData["Num"] = service.Num

	// Marshal the updated data back to JSON
	updatedData, err := json.Marshal(cartData)
	if err != nil {
		// Handle JSON marshal error
		return response.Response{Status: 401, Data: nil, Msg: "Error marshaling JSON", Error: err.Error()}
	}

	// Save the updated data back to the Redis cache
	err = cache.RDB.HSet(cartKey, field, updatedData).Err()
	if err != nil {
		// Handle error when updating cart data
		return response.Response{Status: 401, Data: nil, Msg: "Error updating cart data", Error: err.Error()}
	}
	return response.Response{Status: 201, Data: nil, Msg: "Cart updated successfully", Error: ""}
	//var cart models.Cart
	//
	//err := models.DB.
	//	Where("user_id=? AND product_id=?", service.UserId, service.ProductId).Find(&cart).Error
	//if err != nil {
	//	return response.Response{Status: 401, Data: nil, Msg: "", Error: err.Error()}
	//}
	//
	//cart.Num = service.Num
	//err = models.DB.Save(&cart).Error
	//if err != nil {
	//	return response.Response{Status: 401, Data: nil, Msg: "", Error: err.Error()}
	//}
}
