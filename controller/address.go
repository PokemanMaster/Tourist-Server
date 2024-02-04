package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"qiqi-go/response"
	"qiqi-go/service"
)

// CreateAddress 创建收货地址
func CreateAddress(ctx *gin.Context) {
	services := service.CreateAddressService{}
	err := ctx.ShouldBind(&services)
	if err != nil {
		ctx.JSON(400, response.ErrorResponse(err))
		log.Println("error", err)
	} else {
		res := services.Create()
		ctx.JSON(200, res) // 解析数据JSON
	}
}

// ShowAddress 展示收货地址
func ShowAddress(ctx *gin.Context) {
	services := service.ShowAddressService{}
	res := services.Show(ctx.Param("user_id"))
	ctx.JSON(200, res) // 解析数据JSON
}

// UpdateAddress 修改收货地址
func UpdateAddress(ctx *gin.Context) {
	services := service.UpdateAddressService{}
	err := ctx.ShouldBind(&services)
	if err != nil {
		ctx.JSON(400, response.ErrorResponse(err))
		log.Println("error", err)
	} else {
		res := services.Update()
		ctx.JSON(200, res) // 解析数据JSON
	}
}

// DeleteAddress 删除收货地址
func DeleteAddress(ctx *gin.Context) {
	services := service.DeleteAddressService{}
	err := ctx.ShouldBind(&services)
	if err != nil {
		ctx.JSON(400, response.ErrorResponse(err))
		log.Println("error", err)
	} else {
		res := services.Delete()
		ctx.JSON(200, res) // 解析数据JSON
	}
}
