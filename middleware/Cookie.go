package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// SetCookie 创建cookie
func SetCookie(ctx *gin.Context) {
	// 使用 ctx.SetCookie 方法可以设置一个名为 "username" 的 Cookie，值为 "John Doe"。
	// 您可以根据需要调整其他参数，如过期时间、路径、域名等。
	ctx.SetCookie("username", "John Doe", 3600, "/", "localhost", false, true)
	// 返回响应或进行其他操作
	//ctx.Redirect(http.StatusOK, "/")
	//ctx.String(http.StatusOK, "Cookie set successfully")
}

func GetCookie(ctx *gin.Context) {
	//使用c.Cookie("username")从请求的上下文中获取名为"username"的Cookie的值。
	username, err := ctx.Cookie("Set-Cookie")
	if err != nil {
		// 处理错误情况
		ctx.String(http.StatusNotFound, "Cookie not found")
		return
	}
	ctx.Redirect(200, "/")
	// 使用获取到的cookie值
	ctx.String(http.StatusOK, "Username: "+username)
}

func DeleteCookie(c *gin.Context) {
	c.SetCookie("username", "", -1, "/", "localhost", false, true)
	c.String(http.StatusOK, "Cookie deleted successfully")
}
