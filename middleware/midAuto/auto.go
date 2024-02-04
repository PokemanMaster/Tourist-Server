package midAuto

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"qiqi-go/middleware"
)

// AuthCheck 检测用户是否登录状态
func AuthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		userClaims, err := middleware.AnalyseToken(token)
		if err != nil {
			c.Abort()
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "用户认证不通过",
			})
			return
		}
		c.Set("user_claims", userClaims)
		c.Next()
	}
}
