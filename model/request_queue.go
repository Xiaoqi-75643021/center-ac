package model

import (
	"center-air-conditioning-interactive/constants"
	"fmt"
	"sync"
)

type RequestQueue struct {
	queue []*BlowRequest
	mu    sync.Mutex
}

var instance *RequestQueue
var requestQueueOnce sync.Once

func GetRequestQueue() *RequestQueue {
	requestQueueOnce.Do(func() {
		instance = &RequestQueue{
			queue: make([]*BlowRequest, 0),
		}
	})
	return instance
}

func (rq *RequestQueue) AddRequest(request *BlowRequest) {
	rq.mu.Lock()
	defer rq.mu.Unlock()
	rq.queue = append(rq.queue, request)
	rq.refreshRequestsStatus()

	fmt.Println(rq.queue)
}

func (rq *RequestQueue) GetNextRequest() *BlowRequest {
	rq.mu.Lock()
	defer rq.mu.Unlock()
	if len(rq.queue) == 0 {
		return nil
	}
	nextRequest := rq.queue[0]
	rq.queue = rq.queue[1:]
	return nextRequest
}

func (rq *RequestQueue) RemoveRequestByRoomId(roomId string) {
	rq.mu.Lock()
	defer rq.mu.Unlock()
	for i, req := range rq.queue {
		if req.RoomId == roomId {
			rq.queue = append(rq.queue[:i], rq.queue[i+1:]...)
			break
		}
	}
	rq.refreshRequestsStatus()
}

func (rq *RequestQueue) refreshRequestsStatus() {
	for i, req := range rq.queue {
		if i < 3 {
			req.RequestStatus = constants.RequestStatusDoing
		} else {
			break
		}
	}
}

func (rq *RequestQueue) UpdateRequestStatusByRoomId(roomId string) {
	rq.mu.Lock()
	defer rq.mu.Unlock()
	for _, req := range rq.queue {
		if req.RoomId == roomId {
			req.RequestStatus = constants.RequestStatusDone
			break
		}
	}
}

func (rq *RequestQueue) QueryRequestByRoomId(roomId string) *BlowRequest {
	for _, req := range rq.queue {
		if req.RoomId == roomId {
			return req
		}
	}
	return nil
}