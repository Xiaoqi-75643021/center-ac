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

// func handleRequest(ac *model.CentralAC, request *model.RoomACRequest) {
// 	startTime := time.Now()
// 	roomId := request.RoomAC.RoomId
// 	startTemp := ac.CurrentTemp
// 	fanSpeed := request.RoomAC.FanSpeed

// 	requestInfo := model.RequestInfo{
// 		StartTime: startTime,
// 		StartTemp: startTemp,
// 		FanSpeed:  constants.FanSpeedToString[fanSpeed],
// 	}

// 	ac.Mu.Lock()
// 	roomInfo := ac.Rooms[roomId]
// 	roomInfo.RequestLogs = append(roomInfo.RequestLogs, requestInfo)
// 	ac.Rooms[roomId] = roomInfo
// 	ac.Mu.Unlock()

// 	ticker := time.NewTicker(1 * time.Second)
// 	defer ticker.Stop()

// 	for range ticker.C {
// 		ac.Mu.Lock()
// 		if request.RequestStatus == constants.RequestStatusDone {
// 			ac.Rooms[roomId].RequestLogs[len(ac.Rooms[roomId].RequestLogs)-1].EndTime = time.Now()
// 			ac.Rooms[roomId].RequestLogs[len(ac.Rooms[roomId].RequestLogs)-1].EndTemp = ac.CurrentTemp
// 			ac.Mu.Unlock()
// 			break
// 		}

// 		energyUsed := calculateEnergyUsed(constants.FanSpeedToString[fanSpeed])
// 		ac.Rooms[roomId].RequestLogs[len(ac.Rooms[roomId].RequestLogs)-1].EnergyUsed += energyUsed
// 		ac.Rooms[roomId].RequestLogs[len(ac.Rooms[roomId].RequestLogs)-1].Cost += energyUsed * 5
// 		ac.Mu.Unlock()
// 	}
// }

// func calculateEnergyUsed(fanSpeed string) float64 {
// 	switch fanSpeed {
// 	case "low":
// 		return 0.8
// 	case "medium":
// 		return 1.0
// 	case "high":
// 		return 1.2
// 	default:
// 		return 0
// 	}
// }
