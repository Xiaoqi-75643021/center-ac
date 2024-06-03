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

	request := requestQueue.QueryRequestByRoomId(roomId)
	if request != nil {
		requestQueue.RemoveRequestByRoomId(roomId)
	}

	requestQueue.AddRequest(&model.BlowRequest{
		RoomId: roomId,
		StartTime: time.Now(),
		RequestStatus: constants.RequestStatusPending,
	})

	room.RoomAC.TargetTemp = targetTemp
	room.RoomAC.FanSpeed = constants.FanSpeedToInt[fanSpeed]

	message := fmt.Sprintf("Blowing started | Target Temperature: %v°C | Fanspeed:%v", targetTemp, fanSpeed)
	model.GetPrinterInstance().Print("blow", roomId, message)

	go updateEnergyAndCost(request)
	return nil
}

func StopBlowing(roomId string) {
	requestQueue := model.GetRequestQueue()
	requestQueue.UpdateRequestStatusByRoomId(roomId)

	model.GetPrinterInstance().Print("blow", roomId, "Blowing stopped")
}