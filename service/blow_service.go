package service

import (
	"center-air-conditioning-interactive/constants"
	"center-air-conditioning-interactive/model"
	"errors"
	"fmt"
	"time"
)

func StartBlowing(roomId string, targetTemp float64, fanSpeed string) error {
	requestQueue := model.GetRequestQueue()
	rm := model.GetRoomManagerInstance()

	room, exists := rm.Rooms[roomId]
	if !exists {
		return errors.New("房间不存在")
	}

	requestQueue.AddRequest(&model.BlowRequest{
		RoomId: roomId,
		StartTime: time.Now(),
		RequestStatus: constants.RequestStatusPending,
		EnergyUsed: make([]float64, 3),
	})

	ac := model.GetCentralACInstance()
	ac.SetStatus(constants.CentralStatusActive)

	request := requestQueue.QueryRequestByRoomId(roomId)

	room.RoomAC.TargetTemp = targetTemp
	room.RoomAC.FanSpeed = constants.FanSpeedToInt[fanSpeed]
	room.RoomAC.SwitchTime++

	message := fmt.Sprintf("Blowing started | Target Temperature: %.1f°C | Fanspeed:%v", targetTemp, fanSpeed)
	model.GetPrinterInstance().Print("blow", roomId, message)

	go CalculateEnergyAndCost(request)
	return nil
}

func StopBlowing(roomId string) {
	requestQueue := model.GetRequestQueue()
	requestQueue.UpdateRequestStatusByRoomId(roomId)
	if requestQueue.IsEmpty() {
		ac := model.GetCentralACInstance()
		ac.SetStatus(constants.CentralStatusStandBy)
	}

	model.GetPrinterInstance().Print("blow", roomId, "Blowing stopped")
}