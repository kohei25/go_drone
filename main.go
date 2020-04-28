package main

import (
	"log"

	"github.com/nicodot25/go_drone.git/config"
	"github.com/nicodot25/go_drone.git/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	log.Println("test")
	// droneManager := models.NewDronManager()
	// droneManager.TakeOff()
	// time.Sleep(10*time.Second)
	// droneManager.Land()
}
