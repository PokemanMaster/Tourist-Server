package api

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"qiqi-go/controller"
	"qiqi-go/middleware"
	"qiqi-go/middleware/midSDK"
)

func CollectRoute() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.CORS())                      // 同源请求
	store := cookie.NewStore([]byte(midSDK.VERSION))   // 创建一个基于 cookie 的会话存储
	router.Use(sessions.Sessions("my_session", store)) // 启用会话管理。

	// 未登录用户
	v1 := router.Group("/api/v1")
	{
		v1.POST("user/register", controller.UserRegister)      // 用户注册
		v1.POST("user/login", controller.UserLogin)            // 用户登录
		v1.GET("user/category", controller.UserCategoryImages) // 验证码图片

		v1.GET("home/product", controller.HomeProducts)  // 商品首页
		v1.GET("home/carousel", controller.HomeCarousel) // 首页的轮播图
		v1.GET("rankings", controller.ListProductsRank)  // 获取排行榜前10的商品

		v1.POST("sort/product", controller.SortProducts)      // 商品分类
		v1.POST("sort/search", controller.SortProductsSearch) // 商品搜索
		v1.GET("product/:id", controller.ShowProducts)        // 商品详情页

		// 已登录用户
		authed := v1.Group("/")
		authed.Use(middleware.Token())
		{
			authed.GET("ping", controller.UserToken)      // 验证用户token是否正确
			authed.POST("user/info", controller.UserInfo) // 获取用户信息
			authed.POST("user/edit", controller.UserEdit) // 用户修改信息操作

			authed.POST("cart", controller.CreateCart)   // 创建购物车
			authed.GET("cart/:id", controller.ShowCart)  // 展示购物车
			authed.PUT("cart", controller.UpdateCart)    // 修改购物车
			authed.DELETE("cart", controller.DeleteCart) // 删除购物车

			authed.POST("favorites", controller.CreateFavorites)   // 创建收藏夹
			authed.GET("favorites/:id", controller.ShowFavorites)  // 展示收藏夹
			authed.DELETE("favorites", controller.DeleteFavorites) // 删除收藏夹

			authed.POST("address", controller.CreateAddress)   // 创建收货地址
			authed.GET("address/:id", controller.ShowAddress)  // 展示收货地址
			authed.PUT("address", controller.UpdateAddress)    // 修改收货地址
			authed.DELETE("address", controller.DeleteAddress) // 删除收货地址

			authed.POST("order", controller.CreateOrder)   // 创建订单
			authed.GET("order/:num", controller.ShowOrder) // 展示订单
		}
	}
	//管理员
	v2 := router.Group("/api/v2")
	{
		v2.POST("/admin/register", controller.AdminRegister) // 管理员注册
		v2.POST("/admin/login", controller.AdminLogin)       // 管理员注册
	}

	return router
}
