package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"qiqi-go/response"
	"qiqi-go/service"
)

// HomeProducts 首页商品 home/products
func HomeProducts(ctx *gin.Context) {
	services := service.ListProductHomeService{}
	err := ctx.ShouldBind(&services)
	if err != nil {
		ctx.JSON(400, response.ErrorResponse(err))
		log.Println("error", err)
	} else {
		res := services.HomeList()
		ctx.JSON(200, res) // 解析数据JSON
	}
}

// HomeCarousel 首页轮播图 home/carousel
func HomeCarousel(ctx *gin.Context) {
	services := service.ShowCarouselsService{}
	err := ctx.ShouldBind(&services)
	if err != nil {
		ctx.JSON(400, response.ErrorResponse(err))
		log.Println("error", err)
	} else {
		res := services.GetCarouselsImages()
		ctx.JSON(200, res) // 解析数据JSON
	}
}

// ListProductsRank 商品排行榜
func ListProductsRank(ctx *gin.Context) {
	services := service.ListProductsRankService{}
	res := services.ListRank()
	ctx.JSON(200, res)
}
