package main

import (
	"sync"

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

	wg := sync.WaitGroup{}

	wg.Add(1)

	go compass.Start()
	go motor.Start(&wg)
	go steering.Start(&compass)

	wg.Wait()
}
