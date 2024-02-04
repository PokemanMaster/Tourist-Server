package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func indexHandler(c *gin.Context) {
	session := sessions.Default(c)
	// 检查会话是否存在
	username := session.Get("username")
	if username == nil {
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}
	// 用户已登录，显示欢迎信息
	c.HTML(http.StatusOK, "index.html", gin.H{
		"username": username,
	})
}

func sessionHandler(c *gin.Context) {
	session := sessions.Default(c)
	// 检查会话是否存在
	username := session.Get("username")
	if username == nil {
		c.Redirect(http.StatusSeeOther, "/api/user/login")
		return
	}
	// 用户已登录，显示仪表盘
	c.HTML(http.StatusOK, "dashboard.html", gin.H{
		"username": username,
	})
}

func logoutHandler(c *gin.Context) {
	session := sessions.Default(c)
	// 销毁会话
	session.Clear()
	session.Save()
	c.Redirect(http.StatusSeeOther, "/login")
}
