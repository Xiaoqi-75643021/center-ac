package service

import (
	"center-air-conditioning-interactive/model"
	"center-air-conditioning-interactive/pkg/jwt"
	"errors"
	"log"
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

	
	log.Printf("[Room%v]: Login in successfully\n", roomId)
	return token, nil
}