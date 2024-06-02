package handler

// import (
// 	"center-air-conditioning-interactive/constants"
// 	"center-air-conditioning-interactive/model"
// 	"center-air-conditioning-interactive/service"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func GetStatus(c *gin.Context) {
// 	ac := model.GetCentralACInstance()
// 	Respond(c, http.StatusOK, 0, "获取状态成功", gin.H{
// 		"mode":        constants.CentralACModeToString[ac.Mode],
// 		"temperature": ac.CurrentTemp,
// 		"status":      constants.CentralACStatusToString[ac.Status],
// 	})
// }

// func SetMode(c *gin.Context) {
// 	var request struct {
// 		Mode string `json:"mode" binding:"required"`
// 	}
// 	if err := c.ShouldBindJSON(&request); err != nil {
// 		Respond(c, http.StatusBadRequest, 1, err.Error(), nil)
// 		return
// 	}

// 	ac := model.GetCentralACInstance()
// 	ac.Mu.Lock()
// 	defer ac.Mu.Unlock()

// 	if request.Mode == "cooling" || request.Mode == "heating" {
// 		ac.Mode = request.Mode
// 		if ac.Mode == "cooling" {
// 			ac.CurrentTemp = ac.DefaultCoolingTemp
// 		} else {
// 			ac.CurrentTemp = ac.DefaultHeatingTemp
// 		}
// 		Respond(c, http.StatusOK, 0, "模式设置成功", nil)
// 	} else {
// 		Respond(c, http.StatusBadRequest, 1, "无效的模式", nil)
// 	}
// }

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
