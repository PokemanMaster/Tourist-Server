package service

import (
	"fmt"
	"github.com/go-redis/redis"
	"math/rand"
	"os"
	"qiqi-go/cache"
	"qiqi-go/models"
	"qiqi-go/response"
	"strconv"
	"time"
)

// CreateOrderService 前端需要传入参数
type CreateOrderService struct {
	AddressID uint `json:"address_id"` // 地址id
	UserID    uint `json:"user_id"`    // 用户id
	ProductID uint `json:"product_id"` // 商品id
	Num       uint `json:"num"`        // 商品数量
}

// Create 用户创建一个订单
func (service *CreateOrderService) Create() response.Response {
	//查找用户的地址
	address := models.Address{}
	err := models.DB.First(&address, service.AddressID).Error
	if err != nil {
		return response.Response{Status: 402, Data: nil, Msg: "", Error: err.Error()}
	}

	// 存储订单内容
	order := models.Orders{
		UserID:    service.UserID,
		ProductID: service.ProductID,
		Num:       service.Num,
		Type:      1, // 状态
	}
	order.AddressName = address.Name
	order.AddressPhone = address.Phone
	order.Address = address.Address

	//生成随机订单号 = 随机值 + 商品ID + 用户ID
	number := fmt.Sprintf("%09v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000000))
	productNum := strconv.Itoa(int(service.ProductID))
	userNum := strconv.Itoa(int(service.UserID))
	number = number + productNum + userNum
	orderNum, err := strconv.ParseUint(number, 10, 64)
	if err != nil {
		return response.Response{Status: 401, Data: nil, Msg: "", Error: err.Error()}
	}
	order.OrderNum = orderNum

	// 把订单数据存入数据库
	err = models.DB.Create(&order).Error
	if err == nil {
		return response.Response{Status: 402, Data: nil, Msg: "", Error: err.Error()}
	}

	// 把订单放入redis，并设置15分钟
	data := redis.Z{Score: float64(time.Now().Unix()) + 15*time.Minute.Seconds(), Member: orderNum}
	cache.RDB.ZAdd(os.Getenv("REDIS_ZSET_KEY"), data)
	return response.Response{Status: 201, Data: nil, Msg: "", Error: ""}
}
