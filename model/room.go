package model

import (
	"sync"
	"time"
)

type Room struct {
	RoomId      string
	Identity    string
	CurrentTemp float64
	RoomAC      *RoomAC
}

type RoomManager struct {
	Rooms map[string]Room

	mu sync.Mutex
}

type RoomAC struct {
	TargetTemp float64
	FanSpeed   int // low/1, medium/2, or high/3

	Status int // on/1 or off/0

	SwitchTime   int
	CostTracker  *CostTracker
	BlowRequests []*BlowRequest
}

type BlowRequest struct {
	RoomId        string
	StartTime     time.Time
	EndTime       time.Time
	StartTemp     float64
	EndTemp       float64
	EnergyUsed    []float64 // high, medium, low  cost
	Cost          float64
	RequestStatus int // Pending/1, Doing/2, Done/3
}

var RoomManagerInstance *RoomManager
var roomOnce sync.Once

func GetRoomManagerInstance() *RoomManager {
	roomOnce.Do(func() {
		RoomManagerInstance = &RoomManager{
			Rooms: make(map[string]Room),
		}
	})
	return RoomManagerInstance
}

func (rm *RoomManager) AddRoom(room Room) {
	rm.mu.Lock()
	defer rm.mu.Unlock()
	rm.Rooms[room.RoomId] = room
}

func (rm *RoomManager) RemoveRoom(roomId string) {
	rm.mu.Lock()
	defer rm.mu.Unlock()
	delete(rm.Rooms, roomId)
}
