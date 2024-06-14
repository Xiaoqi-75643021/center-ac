package middleware

import (
	"center-air-conditioning-interactive/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

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