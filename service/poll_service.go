package service

import (
	"center-air-conditioning-interactive/constants"
	"center-air-conditioning-interactive/model"
	"errors"
	"fmt"
)

func QueryEnergyAndCostByRoomId(roomId string) (float64, float64, error) {
	var energy, amount float64
	requestQueue := model.GetRequestQueue()
	request := requestQueue.QueryRequestByRoomId(roomId)
	if request == nil {
		return energy, amount, errors.New("房间无送风请求")
	}

	for _, v := range request.EnergyUsed {
		energy += v
	}
	amount = request.Cost

	return energy, amount, nil
}

func QueryCentralACMode() (string, error) {
	ac := model.GetCentralACInstance()
	mode := constants.CentralACModeToString[ac.Mode]
	return mode, nil
}

func QueryRefreshRate() (int, error) {
	ac := model.GetCentralACInstance()
	refreshRate := ac.RefreshRate
	return refreshRate, nil
}

func UpdateRoomByRoomId(roomId string, temperature float64, status int) error {
	rm := model.GetRoomManagerInstance()
	room, exists := rm.Rooms[roomId];
	if !exists {
		return errors.New("房间不存在")
	}
	room.CurrentTemp = temperature
	room.RoomAC.Status = status

	message := fmt.Sprintf("Temperature:%v°C | Status:%v", temperature, constants.RoomStatusToString[status])
	model.GetPrinterInstance().Print("statusReport", roomId, message)

	return nil
}

func QueryBlowRequestStatusByRoomId(roomId string) (string, error) {
	requestQueue := model.GetRequestQueue()

	request := requestQueue.QueryRequestByRoomId(roomId)
	if request == nil {
		return "", errors.New("房间无送风请求")
	}
	requestStatus := request.RequestStatus

	return constants.RequestStatusToString[requestStatus], nil
}
