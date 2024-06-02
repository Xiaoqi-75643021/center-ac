package model

import (
	"sync"
	"time"
)

type Timer struct {
	ticker *time.Ticker
	done   chan bool
}

var TimerInstance *Timer
var timerOnce sync.Once

func GetTimerInstance() *Timer {
	timerOnce.Do(func() {
		TimerInstance = &Timer{
			ticker: time.NewTicker(24 * time.Hour),
			done:   make(chan bool),
		}
	})
	return TimerInstance
}

func (t *Timer) Start() {
	go func() {
		for {
			select {
			case <-t.done:
				return
			case <-t.ticker.C:
				t.updateCostTrackers()
			}
		}
	}()
}

func (t *Timer) Stop() {
	t.ticker.Stop()
	t.done <- true
}

func (t *Timer) updateCostTrackers() {
	roomManager := GetRoomManagerInstance()
	roomManager.mu.Lock()
	defer roomManager.mu.Unlock()

	for _, room := range roomManager.Rooms {
		if room.RoomAC != nil && room.RoomAC.CostTracker != nil {
			totalCost := 0.0
			now := time.Now()
			for _, request := range room.RoomAC.BlowRequests {
				if request.EndTime.After(now.Add(-24 * time.Hour)) && request.EndTime.Before(now) {
					totalCost += request.Cost
				}
			}
			room.RoomAC.CostTracker.AddCost(totalCost)
		}
	}
}
