package models

import (
	"github.com/shopspring/decimal"
	"qiqi-go/cache"
	"strconv"
)

// Products 商品详情表
type Products struct {
	ProductsId          int             // 商品唯一标识
	ProductsImage       string          // 商品图片
	ProductsName        string          // 商品名字
	ProductsLabel       string          // 商品标签
	ProductsTitle       string          // 商品标题，用于展示一级详情页
	ProductsDescription string          // 商品描述，用于展示二级详情页
	ProductsMoney       decimal.Decimal // 商品金额
}

// AddView 商品详情预览
func (product *Products) AddView() {
	// 增加商品的点击数
	cache.RDB.Incr(cache.ProductViewKey(product.ProductsId))
	// 增加商品排行榜rank的点击数
	cache.RDB.ZIncrBy(cache.RankKey, 1, strconv.Itoa(product.ProductsId))
}

// View 获取点击数
//func (product *Products) View() uint64 {
//	countStr, _ := common.RDB.Get(common.ProductViewKey(product.ProductsId)).Result()
//	count, _ := strconv.ParseUint(countStr, 10, 64)
//	return count
//}

//// AddElecRank 增加家电排行点击数
//func (product *Products) AddElecRank() {
//	// 增加家电排行点击数
//	common.RDB.ZIncrBy(common.ElectricalRank, 1, strconv.Itoa(product.ProductsId))
//}
//
//// AddAcceRank 增加配件排行点击数
//func (product *Products) AddAcceRank() {
//	// 增加配件排行点击数
//	common.RDB.ZIncrBy(common.AccessoryRank, 1, strconv.Itoa(product.ProductsId))
//}
