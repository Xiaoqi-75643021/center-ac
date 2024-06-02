package router

import (
	"center-air-conditioning-interactive/handler"
	"center-air-conditioning-interactive/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
    authGroup := r.Group("/auth")
    {
        authGroup.POST("/login", handler.Login)
    }

    roomGroup := r.Group("/room")
    roomGroup.Use(middleware.AuthMiddleware())
    {
        roomGroup.POST("/logout", handler.Logout)
        blowGroup := roomGroup.Group("/blowing")
        {
            blowGroup.POST("/start", handler.StartBlowing)
            blowGroup.POST("/stop", handler.StopBlowing)
        }
    
        pollGroup := roomGroup.Group("/poll")
        {
            pollGroup.POST("/request", handler.QueryBlowRequestStatus)
            pollGroup.POST("/billing", handler.QueryBilling)
            pollGroup.POST("/room_status", handler.UpdateRoomStatus)
        }
    
        roomGroup.POST("/report", handler.ReportForms)
    }
}
