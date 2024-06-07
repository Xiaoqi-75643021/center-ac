package handler

import (
	"center-air-conditioning-interactive/constants"
	"center-air-conditioning-interactive/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetMode(c *gin.Context) {
	type request struct {
		Mode string `json:"mode" binding:"required"`
	}
	var req request
	if err := c.ShouldBindJSON(&req); err != nil {
		Respond(c, http.StatusBadRequest, 1, err.Error(), nil)
		return
	}

	ac := model.GetCentralACInstance()

	switch req.Mode {
	case "Cool":
		ac.Mode = constants.CoolMode
		ac.DefaultTemp = constants.DefaultCoolingTemp
	case "Warm":
		ac.Mode = constants.HeatMode
		ac.DefaultTemp = constants.DefaultHeatingTemp
	default:
		Respond(c, http.StatusBadRequest, 1, "无效的模式", nil)
		return
	}

	Respond(c, http.StatusOK, 0, "中央空调模式设置成功", nil)
}

// func SetTemperature(c *gin.Context) {
// 	var json struct {
// 		Temperature float64 `json:"temperature" binding:"required"`
// 	}
// 	if err := c.ShouldBindJSON(&json); err != nil {
// 		Respond(c, http.StatusBadRequest, 1, err.Error(), nil)
// 		return
// 	}

// 	ac := model.GetCentralACInstance()
// 	ac.Mu.Lock()
// 	defer ac.Mu.Unlock()

// 	if (ac.Mode == "cooling" && json.Temperature >= 18.0 && json.Temperature <= 25.0) ||
// 		(ac.Mode == "heating" && json.Temperature >= 25.0 && json.Temperature <= 30.0) {
// 		ac.CurrentTemp = json.Temperature
// 		Respond(c, http.StatusOK, 0, "温度设置成功", nil)
// 	} else {
// 		Respond(c, http.StatusBadRequest, 1, "温度超出范围", nil)
// 	}
// }

// func SendRequest(c *gin.Context) {
// 	var json struct {
// 		RoomID      string  `json:"room_id" binding:"required"`
// 		Temperature float64 `json:"temperature" binding:"required"`
// 		FanSpeed    string  `json:"fan_speed" binding:"required"`
// 	}
// 	if err := c.ShouldBindJSON(&json); err != nil {
// 		Respond(c, http.StatusBadRequest, 1, err.Error(), nil)
// 		return
// 	}

// 	if json.FanSpeed != "low" && json.FanSpeed != "medium" && json.FanSpeed != "high" {
// 		Respond(c, http.StatusBadRequest, 1, "无效的风速", nil)
// 		return
// 	}

// 	if err := service.HandleRequest(json.RoomID, json.Temperature, json.FanSpeed); err != nil {
// 		Respond(c, http.StatusInternalServerError, 1, err.Error(), nil)
// 		return
// 	}
// 	Respond(c, http.StatusOK, 0, "请求处理成功", nil)
// }

// func GetBilling(c *gin.Context) {
// 	ac := model.GetCentralACInstance()
// 	Respond(c, http.StatusOK, 0, "获取账单成功", ac.Billing)
// }

// func GetReport(c *gin.Context) {
// 	roomID := c.Query("room_id")
// 	report, err := service.GenerateReport(roomID)
// 	if err != nil {
// 		Respond(c, http.StatusInternalServerError, 1, err.Error(), nil)
// 		return
// 	}
// 	Respond(c, http.StatusOK, 0, "获取报告成功", report)
// }
