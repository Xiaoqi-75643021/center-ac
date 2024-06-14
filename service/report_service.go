package service

import (
	"center-air-conditioning-interactive/model"
	"errors"
	"fmt"
	"log"
	"time"
)

func QueryRoomLogByRoomId(roomId string, period string) (int, []*model.BlowRequest, float64, error) {
	rm := model.GetRoomManagerInstance()
	
	room, exists := rm.Rooms[roomId]
	if !exists {
		return 0, nil, 0, errors.New("房间不存在") 
	}

	var switchTime int
	var requests []*model.BlowRequest
	var totalCost float64

	switchTime = room.RoomAC.SwitchTime
	requests = room.RoomAC.BlowRequests
	switch period {
	case "daily":
		totalCost = room.RoomAC.CostTracker.GetDayTotal()
	case "weekly":
		totalCost = room.RoomAC.CostTracker.GetWeekTotal()
	case "monthly":
		totalCost = room.RoomAC.CostTracker.GetMonthTotal()
	default:
		return 0, nil, 0, errors.New("无效的时间段")
	}

	return switchTime, requests, totalCost, nil
}

func ExportRoomReport(roomId string, period string) error {
	rm := model.GetRoomManagerInstance()
	
	room, exists := rm.Rooms[roomId]
	if !exists {
		return errors.New("房间不存在") 
	}

	var switchTime int
	var requests []*model.BlowRequest
	var totalCost float64

	switchTime = room.RoomAC.SwitchTime
	requests = room.RoomAC.BlowRequests
	switch period {
	case "daily":
		totalCost = room.RoomAC.CostTracker.GetDayTotal()
	case "weekly":
		totalCost = room.RoomAC.CostTracker.GetWeekTotal()
	case "monthly":
		totalCost = room.RoomAC.CostTracker.GetMonthTotal()
	default:
		return errors.New("无效的时间段")
	}

	// 记录到日志文件
	logMessage := createLogMessage(roomId, switchTime, requests, totalCost, period)
	log.Println(logMessage)

	return nil
}

func createLogMessage(roomId string, switchTime int, requests []*model.BlowRequest, totalCost float64, period string) string {
	// 创建详细的日志信息
	requestsInfo := ""
	for _, request := range requests {
		requestsInfo += request.String() + "; "
	}
	return time.Now().Format("2006-01-02 15:04:05") + " - RoomID: " + roomId + ", SwitchTime: " + string(rune(switchTime)) + ", Requests: [" + requestsInfo + "], TotalCost: " + fmt.Sprintf("%.2f", totalCost) + ", Period: " + period
}