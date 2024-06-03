package service

import (
	"center-air-conditioning-interactive/constants"
	"center-air-conditioning-interactive/model"
	"fmt"

	"time"
)

func updateEnergyAndCost(request *model.BlowRequest) {

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		if request.RequestStatus == constants.RequestStatusDone {
			UpdateRoomRequestLog(request)
			return
		}

		request.EnergyUsed[0] += constants.LowSpeedConsumedPerSecond
		request.EnergyUsed[1] += constants.MediumSpeedConsumedPerSecond
		request.EnergyUsed[2] += constants.HighSpeedConsumedPerSecond
		request.Cost = constants.CostPerEnergy * (request.EnergyUsed[0] + request.EnergyUsed[1] + request.EnergyUsed[2])

		message := fmt.Sprintf("Energy Used: %v | Cost: %v", request.EnergyUsed, request.Cost)
		model.GetPrinterInstance().Print("pollReport", request.RoomId, message)
	}
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
