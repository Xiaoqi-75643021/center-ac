package service

import (
	"center-air-conditioning-interactive/model"
	"center-air-conditioning-interactive/pkg/jwt"
	"errors"
)

func Login(roomId, identity string) (string, error) {
	rm := model.GetRoomManagerInstance()
	if rm.Rooms[roomId].Identity != identity {
		return "", errors.New("房间号或身份证号错误")
	}

	token, err := jwt.GenerateToken(roomId)
	if err != nil {
		return "", err
	}

	model.GetPrinterInstance().Print("auth", roomId, "Login in successfully")
	return token, nil
}