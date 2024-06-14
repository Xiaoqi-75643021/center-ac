package service

import (
	"center-air-conditioning-interactive/model"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func ExportRoomReport(roomId string, period string, savePath string) error {
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

	// 记录到指定文件
	logMessage := createLogMessage(roomId, switchTime, requests, totalCost, period)
	err := saveLogMessage(savePath, logMessage)
	if err != nil {
		return err
	}

	return nil
}

func createLogMessage(roomId string, switchTime int, requests []*model.BlowRequest, totalCost float64, period string) string {
	var builder strings.Builder
	builder.WriteString(time.Now().Format("2006-01-02 15:04:05"))
	builder.WriteString(" - RoomID: " + roomId + "\n")
	builder.WriteString("SwitchTime: " + strconv.Itoa(switchTime) + "\n")
	builder.WriteString("Requests: [")

	for _, request := range requests {
		builder.WriteString("\n")
		builder.WriteString(request.String())
	}

	builder.WriteString("\n]\n")
	builder.WriteString(fmt.Sprintf("TotalCost: %.2f\n", totalCost))
	builder.WriteString("Period: " + period + "\n")

	return builder.String()
}

func saveLogMessage(filePath string, message string) error {
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(message + "\n")
	if err != nil {
		return err
	}

	return nil
}