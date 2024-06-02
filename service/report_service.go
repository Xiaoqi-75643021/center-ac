package service

import (
	"center-air-conditioning-interactive/model"
	"errors"
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
