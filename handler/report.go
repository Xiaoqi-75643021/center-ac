package handler

import (
	"net/http"

	"center-air-conditioning-interactive/service"
	"github.com/gin-gonic/gin"
)

func ReportForms(c *gin.Context) {
	roomId, _ := c.Get("roomId")
	type request struct {
		Period string `json:"period" binding:"required"`
	}

	var req request
	if err := c.ShouldBindJSON(&req); err != nil {
		Respond(c, http.StatusBadRequest, 1, "请求参数错误", nil)
		return
	}

	switchTime, requests, totalCost, err := service.QueryRoomLogByRoomId(roomId.(string), req.Period)
	if err != nil {
		Respond(c, http.StatusInternalServerError, 2, "报表获取失败", gin.H{"error": err.Error()})
		return
	}
	Respond(c, http.StatusOK, 0, "房间"+roomId.(string)+"报表获取成功", gin.H{
		"roomId": roomId,
		"switchTime": switchTime,
		"requests": requests,
		"totalCost": totalCost,
	})
}
