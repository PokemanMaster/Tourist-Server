package service

import (
	"fmt"
	"qiqi-go/cache"
	"qiqi-go/response"
)

// ShowCartService 前端传递过来的数据
type ShowCartService struct{}

// Show 获取购物车的商品
func (service *ShowCartService) Show(userid string) response.Response {
	cartKey := fmt.Sprintf("%s%s", "cart:", userid)
	cartData, err := cache.RDB.HGetAll(cartKey).Result()
	if err != nil {
		return response.Response{Status: 501, Data: nil, Msg: "获取数据失败", Error: err.Error()}
	}

	return response.Response{Status: 201, Data: cartData, Msg: "", Error: ""}
	//id, _ := strconv.ParseUint(userid, 10, 64)
	//var products []struct {
	//	ProductsId          int             // 商品唯一标识
	//	ProductsImage       string          // 商品图片
	//	ProductsName        string          // 商品名字
	//	ProductsLabel       string          // 商品标签
	//	ProductsTitle       string          // 商品标题，用于展示一级详情页
	//	ProductsDescription string          // 商品描述，用于展示二级详情页
	//	ProductsMoney       decimal.Decimal // 商品金额
	//	Num                 uint
	//}
	//
	//err = models.DB.Table("products").
	//	Select("products.products_id, products.products_image, products.products_name, products.products_title , cart.num , products.products_money").
	//	Joins("JOIN cart ON products.products_id = cart.product_id").
	//	Where("cart.user_id = ?", id).
	//	Scan(&products).Error
	//
	//if err != nil {
	//	return response.Response{Status: 501, Data: nil, Msg: "获取数据失败", Error: err.Error()}
	//}
}
