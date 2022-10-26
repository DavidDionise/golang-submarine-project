package motor

import (
	"time"

	"com.pi/submarine/utils"
	"github.com/cgxeiji/servo"
	log "github.com/sirupsen/logrus"
	"github.com/tkanos/gonfig"
)

type Motor struct{}

func (m Motor) Start() {
	utils.WaitGroup.Add(1)
	defer servo.Close()

	config := newMotorConfig()

	motorServo := servo.New(config.Pin)
	motorServo.SetPosition(float64(config.InitialSpeed))

	err := motorServo.Connect()
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(5 * time.Second)

	motorServo.MoveTo(float64(config.InitialSpeed)).Wait()

	time.Sleep(time.Duration(config.RunDurationSeconds) * 1000 * 1000000)

	motorServo.MoveTo(90).Wait()

	utils.WaitGroup.Done()
}

type motorConfig struct {
	Pin                int
	InitialSpeed       int
	RunDurationSeconds int
}

func newMotorConfig() motorConfig {
	config := &motorConfig{}

	err := gonfig.GetConf("motor-config.json", config)

	if err != nil {
		log.Fatal(err)
	}

	return *config
}
