package handler

import (
	"center-air-conditioning-interactive/constants"
	"center-air-conditioning-interactive/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func QueryBilling(c *gin.Context) {
	roomId, _ := c.Get("roomId")

	energy, amount, err := service.QueryEnergyAndCostByRoomId(roomId.(string))
	if err != nil {
		Respond(c, http.StatusInternalServerError, 2, "计费信息获取失败", gin.H{"error": err.Error()})
		return
	}

	message := fmt.Sprintf("房间%v计费信息获取成功", roomId)
	Respond(c, http.StatusOK, 0, message, gin.H{
		"energyConsumed": energy,
		"amountDue":      amount,
	})
}

func UpdateRoomStatus(c *gin.Context) {
	roomId, _ := c.Get("roomId")
	type request struct {
		Temperature float64 `json:"temperature" binding:"required"`
		Status      string  `json:"status" binding:"required"`
		FanSpeed    string  `json:"fan_speed" binding:"required"`
	}
	var req request
	if err := c.ShouldBindJSON(&req); err != nil {
		Respond(c, http.StatusBadRequest, 1, "请求参数错误", nil)
		return
	}

	var statusNum, fanSpeedNum int
	if req.Status == "Warm" {
		statusNum = constants.RoomStatusWarm
	} else if req.Status == "Cool" {
		statusNum = constants.RoomStatusCool
	} else {
		Respond(c, http.StatusBadRequest, 1, "状态参数错误", nil)
		return
	}

	if !(req.FanSpeed == "High" || req.FanSpeed == "Medium" || req.FanSpeed == "Low") {
		Respond(c, http.StatusBadRequest, 1, "风速参数错误", nil)
		return
	} else {
		fanSpeedNum = constants.FanSpeedToInt[req.FanSpeed]
	}

	if err := service.UpdateRoomByRoomId(roomId.(string), req.Temperature, statusNum, fanSpeedNum); err != nil {
		Respond(c, http.StatusInternalServerError, 1, "房间信息同步失败", gin.H{"error": err.Error()})
		return
	}

	mode, _ := service.QueryCentralACMode()
	refreshRate, _ := service.QueryRefreshRate()
	energy, cost := service.CalculateDailyEnergyAndCostByRoomId(roomId.(string))

	message := fmt.Sprintf("房间%v信息同步成功", roomId)
	Respond(c, http.StatusOK, 0, message, gin.H{
		"mode":        mode,
		"refreshRate": refreshRate,
		"daily_statistics": gin.H{
			"energy": energy,
			"cost":   cost,
		},
	})
}

func QueryBlowRequestStatus(c *gin.Context) {
	roomId, _ := c.Get("roomId")

	blowRequestStatus, err := service.QueryBlowRequestStatusByRoomId(roomId.(string))

	if err != nil {
		Respond(c, http.StatusInternalServerError, 1, "送风请求状态获取失败", gin.H{"error": err.Error()})
		return
	}

	message := fmt.Sprintf("房间%v送风请求状态获取成功", roomId)
	Respond(c, http.StatusOK, 0, message, gin.H{
		"requestStatus": blowRequestStatus,
	})
}
