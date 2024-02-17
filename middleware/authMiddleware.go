// 鉴权中间件

package middleware

import (
	"admin-go-api/common/constant"
	"admin-go-api/common/result"
	"github.com/gin-gonic/gin"
	"strings"
)

// AuthMiddleware 鉴权
func AuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			result.Failed(c, int(result.ApiCode.NOAUTH),
				result.ApiCode.GetMessage(result.ApiCode.NOAUTH))
			c.Abort()
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			result.Failed(c, int(result.ApiCode.AUTHFORMATERRROR),
				result.ApiCode.GetMessage(result.ApiCode.AUTHFORMATERRROR))
			c.Abort()
			return
		}
		// todo 校验token
		var token = "token"
		// 存用户信息
		c.Set(constant.ContextKeyUserObj, token)
		c.Next()
	}
}
