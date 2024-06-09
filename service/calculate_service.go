package service

import (
	"center-air-conditioning-interactive/constants"
	"center-air-conditioning-interactive/model"
	"fmt"
	"strings"

	"time"
)

func CalculateEnergyAndCost(request *model.BlowRequest) {

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	rm := model.GetRoomManagerInstance()
	roomAC := rm.Rooms[request.RoomId].RoomAC

	for range ticker.C {
		if request.RequestStatus == constants.RequestStatusDone {
			UpdateRoomRequestLog(request)
			return
		}

		switch roomAC.FanSpeed {
		case constants.FanSpeedLow:
			request.EnergyUsed[0] += constants.LowSpeedConsumedPerSecond
		case constants.FanSpeedMedium:
			request.EnergyUsed[1] += constants.MediumSpeedConsumedPerSecond
		case constants.FanSpeedHigh:
			request.EnergyUsed[2] += constants.HighSpeedConsumedPerSecond
		}
		request.Cost = constants.CostPerEnergy * (request.EnergyUsed[0] + request.EnergyUsed[1] + request.EnergyUsed[2])

		var energys []string
		for _, energy := range request.EnergyUsed {
			energys = append(energys, fmt.Sprintf("%.1f", energy))
		}
		energySliceString := fmt.Sprintf("[%s]", strings.Join(energys, " "))
		message := fmt.Sprintf("Energy Used: %v | Cost: %.1f", energySliceString, request.Cost)
		model.GetPrinterInstance().Print("costReport", request.RoomId, message)
	}
}

func CalculateDailyEnergyAndCostByRoomId(roomId string) (float64, float64) {
	var energy, cost float64
	now := time.Now()
	rm := model.GetRoomManagerInstance()
	roomRequests := rm.Rooms[roomId].RoomAC.BlowRequests
	for _, request := range roomRequests {
		if request.StartTime.After(now) {
			energy += request.EnergyUsed[0] + request.EnergyUsed[1] + request.EnergyUsed[2]
		}
	}

	rq := model.GetRequestQueue()
	request := rq.QueryRequestByRoomId(roomId)
	if request != nil {
		energy += request.EnergyUsed[0] + request.EnergyUsed[1] + request.EnergyUsed[2]
	}
	cost = energy * constants.CostPerEnergy

	return energy, cost
}

func UpdateRoomRequestLog(request *model.BlowRequest) {
	rm := model.GetRoomManagerInstance()
	rq := model.GetRequestQueue()

	request.EndTime = time.Now()
	request.EndTemp = rm.Rooms[request.RoomId].CurrentTemp

	newRequest := &model.BlowRequest{
		RoomId:        request.RoomId,
		StartTime:     request.StartTime,
		EndTime:       request.EndTime,
		StartTemp:     request.StartTemp,
		EndTemp:       request.EndTemp,
		EnergyUsed:    request.EnergyUsed,
		Cost:          request.Cost,
		RequestStatus: request.RequestStatus,
	}
	roomAC := rm.Rooms[request.RoomId].RoomAC
	roomAC.BlowRequests = append(roomAC.BlowRequests, newRequest)
	rq.RemoveRequestByRoomId(request.RoomId)
}
