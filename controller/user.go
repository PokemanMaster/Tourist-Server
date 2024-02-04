package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"qiqi-go/middleware/midSDK"
	"qiqi-go/response"
	"qiqi-go/service"
)

// UserRegister 用户注册接口
func UserRegister(ctx *gin.Context) {
	session := sessions.Default(ctx)
	status := session.Get(midSDK.GEETEST_SERVER_STATUS_SESSION_KEY)
	userID := session.Get("userId")
	println(status, userID)
	services := service.UserRegisterService{}
	err := ctx.ShouldBind(&services)
	if err != nil {
		ctx.JSON(400, response.ErrorResponse(err))
		log.Println("error", err)
	} else {
		res := services.UserRegister(userID, status)
		ctx.JSON(200, res) // 解析数据JSON
	}
}

// UserLogin 用户登录接口
func UserLogin(ctx *gin.Context) {
	session := sessions.Default(ctx)
	status := session.Get(midSDK.GEETEST_SERVER_STATUS_SESSION_KEY)
	userID := session.Get("userId")
	println(status, userID)
	services := service.UserLoginService{}
	err := ctx.ShouldBind(&services)
	if err != nil {
		ctx.JSON(400, response.ErrorResponse(err))
		log.Println("error", err)
	} else {
		res := services.UserLogin(userID, status)
		ctx.JSON(200, res) // 解析数据JSON
	}
}

// UserInfo 获取用户信息接口
func UserInfo(ctx *gin.Context) {
	services := service.UserInfoService{}
	err := ctx.ShouldBind(&services)
	if err != nil {
		ctx.JSON(400, response.ErrorResponse(err))
		log.Println("error", err)
	} else {
		res := services.UserInfo()
		ctx.JSON(200, res) // 解析数据JSON
	}
}

// UserCategoryImages 给用户返回base64码的图片
func UserCategoryImages(ctx *gin.Context) {
	services := service.UserCategoryService{}
	res := services.UserCategoryImages()
	ctx.JSON(200, res)
}

// UserEdit 修改用户信息
func UserEdit(ctx *gin.Context) {
	services := service.UserEditService{}
	err := ctx.ShouldBind(&services)
	if err != nil {
		ctx.JSON(400, response.ErrorResponse(err))
		log.Println("error", err)
	} else {
		res := services.UserEdit()
		ctx.JSON(200, res) // 解析数据JSON
	}
}

// UserToken 检测用户token
func UserToken(ctx *gin.Context) {
	ctx.JSON(200, response.Response{
		Status: 200,
		Msg:    "ok",
	})
}
