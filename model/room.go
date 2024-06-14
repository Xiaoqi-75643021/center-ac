package model

import (
	"fmt"
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
	FanSpeed   int // Low/1, Medium/2, or High/3

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

func (br *BlowRequest) String() string {
	return fmt.Sprintf(
		"RoomId: %s, StartTime: %s, EndTime: %s, StartTemp: %.2f, EndTemp: %.2f, EnergyUsed: %v, Cost: %.2f, RequestStatus: %d",
		br.RoomId,
		br.StartTime.Format("2006-01-02 15:04:05"),
		br.EndTime.Format("2006-01-02 15:04:05"),
		br.StartTemp,
		br.EndTemp,
		br.EnergyUsed,
		br.Cost,
		br.RequestStatus,
	)
}