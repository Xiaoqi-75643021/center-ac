package main

import (
	"center-air-conditioning-interactive/config"
	"center-air-conditioning-interactive/model"
	"center-air-conditioning-interactive/router"
	"center-air-conditioning-interactive/ui"

	"github.com/gin-gonic/gin"
)

func main() {
    config.LoadConfig("config.json")
    r := gin.Default()
    router.SetupRoutes(r)

    timer := model.GetTimerInstance()
	timer.Start()
	defer timer.Stop()

	go ui.RunUI()

    r.Run(":8080")
}
