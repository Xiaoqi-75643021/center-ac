package config

import (
	"center-air-conditioning-interactive/model"

	"encoding/json"
	"io"
	"log"
	"os"
)

type RoomConfig struct {
	RoomId   string `json:"roomId"`
	Identity string `json:"identity"`
}

type Config struct {
	Rooms  []RoomConfig `json:"rooms"`
	JWTKey string       `json:"jwt_secret_key"`
}

var (
	cfg  Config
)

func LoadConfig(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error opening config file: %v", err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	err = json.Unmarshal(bytes, &cfg)
	if err != nil {
		log.Fatalf("Error parsing config file: %v", err)
	}

	initializeRoomManagerInstance(&cfg)
	initializeCentralAC()
	initializeRequestQueue()
}

func initializeCentralAC() {
	model.GetCentralACInstance()
}

func initializeRoomManagerInstance(cfg *Config) {
	roomManager := model.GetRoomManagerInstance()
	for _, room := range cfg.Rooms {
		room := model.Room{
			RoomId:   room.RoomId,
			Identity: room.Identity,
			RoomAC: &model.RoomAC{
				CostTracker: model.NewCostTracker(),
				BlowRequests: make([]*model.BlowRequest, 0),	
			},
		}
		roomManager.AddRoom(room)
		roomManager.Rooms[room.RoomId] = room
	}
}

func initializeRequestQueue() {
	model.GetRequestQueue()
}

func GetJWTSecretKey() string {
	return cfg.JWTKey
}