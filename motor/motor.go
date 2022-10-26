package motor

import (
	"os"
	"sync"
	"time"

	"github.com/cgxeiji/servo"
	log "github.com/sirupsen/logrus"
	"github.com/tkanos/gonfig"
)

type Motor struct{}

var config = newMotorConfig()

func (m Motor) Start(wg *sync.WaitGroup) {
	defer servo.Close()

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

	wg.Done()
}

type motorConfig struct {
	Pin                int
	InitialSpeed       int
	RunDurationSeconds int
}

func newMotorConfig() motorConfig {
	config := &motorConfig{}

	cwd, _ := os.Getwd()

	err := gonfig.GetConf(cwd+"/motor/motor-config.json", config)

	if err != nil {
		log.Fatal(err)
	}

	return *config
}
