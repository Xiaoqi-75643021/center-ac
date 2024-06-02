package handler

import (
	"net/http"

	"center-air-conditioning-interactive/model"
	"github.com/gin-gonic/gin"
)

func AddRoom(c *gin.Context) {
	type request struct {
		RoomId   string  `json:"roomId" binding:"required"`
		Identity string  `json:"identity" binding:"required"`
	}

	var req request
	if err := c.ShouldBindJSON(&req); err != nil {
		Respond(c, http.StatusBadRequest, 1, "请求参数错误", nil)
		return
	}

	rm := model.GetRoomManagerInstance()

	if _, exists := rm.Rooms[req.RoomId]; exists {
		Respond(c, http.StatusConflict, 1, "房间已存在", nil)
		return
	}

	room := model.Room{
		RoomId:      req.RoomId,
		Identity:    req.Identity,
	}
	rm.AddRoom(room)

	Respond(c, http.StatusOK, 0, "房间添加成功", nil)
}

func DeleteRoom(c *gin.Context) {
	type request struct {
		RoomId   string  `json:"roomId" binding:"required"`
		Identity string  `json:"identity" binding:"required"`
	}
	var req request
	if err := c.ShouldBindJSON(&req); err != nil {
		Respond(c, http.StatusBadRequest, 1, "请求参数错误", nil)
		return
	}

	rm := model.GetRoomManagerInstance()

	if _, exists := rm.Rooms[req.RoomId]; !exists {
		Respond(c, http.StatusNotFound, 1, "房间不存在", nil)
		return
	}

	rm.RemoveRoom(req.RoomId)

	Respond(c, http.StatusOK, 0, "房间删除成功", nil)
}
