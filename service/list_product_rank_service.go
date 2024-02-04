package service

import (
	"fmt"
	"qiqi-go/cache"
	"qiqi-go/models"
	"qiqi-go/response"
	"strings"
)

// ListProductsRankService 前端传过来的商品数据
type ListProductsRankService struct{}

// ListRank 展示商品排行榜
func (service *ListProductsRankService) ListRank() response.Response {
	var products []models.Products
	// 从redis读取点击前十的视频  pros:集合数据key
	pros, err := cache.RDB.ZRevRange(cache.RankKey, 0, 9).Result()
	if err != nil {
		return response.Response{Status: 501, Data: nil, Msg: "", Error: err.Error()}
	}
	if len(pros) > 1 {
		order := fmt.Sprintf("FIELD(products_id, %s)", strings.Join(pros, ","))
		err := models.DB.Where("products_id in (?)", pros).Order(order).Find(&products).Error
		if err != nil {
			return response.Response{Status: 502, Data: nil, Msg: "", Error: err.Error()}
		}
	}
	return response.Response{Status: 201, Data: products, Msg: "", Error: ""}
}
