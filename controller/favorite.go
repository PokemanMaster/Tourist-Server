package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"qiqi-go/response"
	"qiqi-go/service"
)

// 收藏夹操作

// CreateFavorites 创建收藏接口
func CreateFavorites(ctx *gin.Context) {
	services := service.CreateFavoriteService{}
	err := ctx.ShouldBind(&services)
	if err != nil {
		ctx.JSON(400, response.ErrorResponse(err))
		log.Println("error", err)
	} else {
		res := services.Create()
		ctx.JSON(200, res) // 解析数据JSON
	}
}

// ShowFavorites 展示收藏夹接口
func ShowFavorites(ctx *gin.Context) {
	services := service.ShowFavoritesService{}
	err := ctx.ShouldBind(&services)
	if err != nil {
		ctx.JSON(400, response.ErrorResponse(err))
		log.Println("error", err)
	} else {
		res := services.Show(ctx.Param("id"))
		ctx.JSON(200, res) // 解析数据JSON
	}
}

// DeleteFavorites 删除收藏夹的接口
func DeleteFavorites(ctx *gin.Context) {
	services := service.DeleteFavoritesService{}
	err := ctx.ShouldBind(&services)
	if err != nil {
		ctx.JSON(400, response.ErrorResponse(err))
		log.Println("error", err)
	} else {
		res := services.Delete()
		ctx.JSON(200, res) // 解析数据JSON
	}
}
