package middleware

import (
	"center-air-conditioning-interactive/handler"
	"center-air-conditioning-interactive/model"
	"center-air-conditioning-interactive/pkg/jwt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			handler.Respond(c, http.StatusUnauthorized, 5, "请求未携带token", nil)
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			handler.Respond(c, http.StatusUnauthorized, 6, "Authorization格式错误", nil)
			c.Abort()
			return
		}

		claims, err := jwt.ParseToken(parts[1])

		if err != nil {
			handler.Respond(c, http.StatusUnauthorized, 7, "token无效", nil)
			c.Abort()
			return
		}

		c.Set("roomId", claims.RoomId)
		
		c.Next()
	}
}

func MasterSwitchMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ac := model.GetCentralACInstance()
		if ac.IsTurnOff() {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "中央空调已关闭，不接收房间请求",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}