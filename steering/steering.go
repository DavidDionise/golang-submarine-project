package steering

import (
	"math"

	"com.pi/submarine/compass"
	"github.com/cgxeiji/servo"
	log "github.com/sirupsen/logrus"
	"github.com/tkanos/gonfig"
)

type Steering struct{}

func (s Steering) Start(compass *compass.Compass) {
	config := newSteeringConfig()
	defer servo.Close()

	steeringServo := servo.New(config.Pin)
	steeringServo.SetPosition(float64(calculateNextAngle(compass.GetAngle(), config.DefaultSetHead)))

	err := steeringServo.Connect()
	if err != nil {
		log.Fatal(err)
	}

	for {
		nextAngle := calculateNextAngle(compass.GetAngle(), config.DefaultSetHead)
		steeringServo.MoveTo(float64(nextAngle))
	}
}

type steeringConfig struct {
	Pin              int
	DefaultSetHead   int
	RateMilliseconds int
	MaxAngle         int
	MinAngle         int
}

func newSteeringConfig() steeringConfig {
	config := &steeringConfig{}
	err := gonfig.GetConf("steering-config.json", config)

	if err != nil {
		log.Fatal(err)
	}

	return *config
}

func calculateNextAngle(currentRead int, setHead int) int {
	config := newSteeringConfig()

	angleOfDeviation := calculateAngleOfDeviation(currentRead, setHead)
	var boundAngle int

	if angleOfDeviation < 0 {
		boundAngle = int(math.Min(float64(config.MinAngle), float64(angleOfDeviation)))
	} else {
		boundAngle = int(math.Max(float64(config.MaxAngle), float64(angleOfDeviation)))
	}

	return boundAngle + 90
}

func calculateAngleOfDeviation(currentRead int, setHead int) int {
	angle := currentRead - setHead

	switch {
	case angle < -180:
		return angle + 360
	case angle > 180:
		return angle - 360
	default:
		return angle
	}
}
