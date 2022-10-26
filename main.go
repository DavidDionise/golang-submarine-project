package main

import (
	"com.pi/submarine/compass"
	"com.pi/submarine/motor"
	"com.pi/submarine/steering"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.DebugLevel)

	compass := compass.Compass{}
	motor := motor.Motor{}
	steering := steering.Steering{}

	go compass.Start()
	go motor.Start()
	go steering.Start(&compass)
}
