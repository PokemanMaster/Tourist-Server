package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"qiqi-go/response"
	"qiqi-go/service"
)

// SortProducts 商品分类接口 sort/products
func SortProducts(ctx *gin.Context) {
	//获取前端请求的数据
	services := service.ListProductSortService{}
	err := ctx.ShouldBind(&services)
	if err != nil {
		ctx.JSON(400, response.ErrorResponse(err))
		log.Println("error", err)
	} else {
		res := services.SortList()
		ctx.JSON(200, res) // 解析数据JSON
	}
}

// SortProductsSearch 分类搜索  sort/search
func SortProductsSearch(ctx *gin.Context) {
	services := service.ListProductSearchService{}
	err := ctx.ShouldBind(&services)
	if err != nil {
		ctx.JSON(400, response.ErrorResponse(err))
		log.Println("error", err)
	} else {
		res := services.SortListSearch()
		ctx.JSON(200, res) // 解析数据JSON
	}
}

// ShowProducts 点击商品进入详情页 products/:id
func ShowProducts(ctx *gin.Context) {
	services := service.ShowProductService{}
	res := services.Show(ctx.Param("id"))
	ctx.JSON(200, res)
}
