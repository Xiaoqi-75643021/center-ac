package handler

import (
	"center-air-conditioning-interactive/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartBlowing(c *gin.Context) {
	type request struct {
		TargetTemp float64 `json:"targetTemp" binding:"required"`
		FanSpeed   string  `json:"fanSpeed" binding:"required"`
	}
	var req request
	if err := c.ShouldBindJSON(&req); err != nil {
		Respond(c, http.StatusBadRequest, 1, "请求参数错误", nil)
		return
	}
	
	roomId, _ := c.Get("roomId")

	service.StartBlowing(roomId.(string), req.TargetTemp, req.FanSpeed)

	message := fmt.Sprintf("房间%v送风开始, 风速：%v", roomId, req.FanSpeed)
	Respond(c, http.StatusOK, 0, message, nil)
}

func StopBlowing(c *gin.Context) {
	roomId, _ := c.Get("roomId")

	service.StopBlowing(roomId.(string))

	message := fmt.Sprintf("房间%v停止送风", roomId)
	Respond(c, http.StatusOK, 0, message, nil)
}