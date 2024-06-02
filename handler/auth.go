package handler

import (
	"center-air-conditioning-interactive/constants"
	"center-air-conditioning-interactive/model"
	"center-air-conditioning-interactive/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	type LoginRequest struct {
		RoomId   string `json:"roomId" binding:"required"`
		Identity string `json:"identity" binding:"required"`
	}

	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		Respond(c, http.StatusBadRequest, 1, "请求参数错误", gin.H{"error": err.Error()})
		return
	}

	token, err := service.Login(req.RoomId, req.Identity)
	if err != nil {
		Respond(c, http.StatusBadRequest, 1, "连接失败", err)
	}

	
	ac := model.GetCentralACInstance()
	Respond(c, http.StatusOK, 0, "房间"+req.RoomId+"连接成功", gin.H{
		"mode":        constants.CentralACModeToString[ac.Mode],
		"defaultTemp": ac.DefaultTemp,
		"refreshRate": ac.RefreshRate,
		"token":       token,
	})
}

func Logout(c *gin.Context) {
	roomId, _ := c.Get("roomId")
	if roomId == "" {
		Respond(c, http.StatusBadRequest, 1, "房间号不能为空", nil)
		return
	}

	Respond(c, http.StatusOK, 0, "房间"+roomId.(string)+"登出成功", nil)
	log.Printf("[Room%v]: Logout successfully\n", roomId)
}
