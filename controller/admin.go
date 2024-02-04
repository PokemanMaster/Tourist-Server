package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"qiqi-go/response"
	"qiqi-go/service"
)

// AdminRegister 管理员注册
func AdminRegister(ctx *gin.Context) {
	services := service.AdminRegisterService{}
	err := ctx.ShouldBind(&services)
	if err != nil {
		ctx.JSON(400, response.ErrorResponse(err))
		log.Println("error", err)
	} else {
		res := services.AdminRegister()
		ctx.JSON(200, res) // 解析数据JSON
	}
}

// AdminLogin 管理员登录
func AdminLogin(ctx *gin.Context) {
	services := service.AdminLoginService{}
	err := ctx.ShouldBind(&services)
	if err != nil {
		ctx.JSON(400, response.ErrorResponse(err))
		log.Println("error", err)
	} else {
		res := services.AdminLogin()
		ctx.JSON(200, res) // 解析数据JSON
	}
}
