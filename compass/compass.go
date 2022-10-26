package compass

import (
	"os"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"

	"com.pi/submarine/utils"
	"github.com/tkanos/gonfig"
	"periph.io/x/conn/v3/i2c"
	"periph.io/x/conn/v3/i2c/i2creg"
	"periph.io/x/host/v3"
)

type Compass struct {
	angle int
}

func (c *Compass) Start() {

	config := newCompassConfig()

	_, err := host.Init()

	if err != nil {
		log.Fatal(err)
	}

	bus, err := i2creg.Open("")

	if err != nil {
		log.Fatal(err)
	}

	addressHex, err := strconv.ParseInt(config.I2cAddress, 0, 16)

	if err != nil {
		log.Fatal(err)
	}

	dev := i2c.Dev{Bus: bus, Addr: uint16(addressHex)}
	write := []byte{0x00, 0x31}

	for {
		buf := make([]byte, 8)
		err = dev.Tx(write, buf)

		if err != nil {
			log.Fatal(err)
		}

		angleValues := [2]byte{buf[1], buf[2]}
		angle := utils.HexToInt(angleValues) / 10

		utils.Mutex.Lock()
		c.angle = angle
		utils.Mutex.Unlock()

		// log.Debug("Current angle: ", c.angle)

		time.Sleep(time.Duration(config.FetchRateMiliseconds * 1000000))
	}
}

func (c Compass) GetAngle() int {
	utils.Mutex.Lock()
	currentAngle := c.angle
	utils.Mutex.Unlock()

	return currentAngle
}

type compassConfig struct {
	I2cAddress           string
	FetchRateMiliseconds int
}

func newCompassConfig() compassConfig {
	config := &compassConfig{}

	cwd, _ := os.Getwd()

	err := gonfig.GetConf(cwd+"/compass/compass-config.json", config)

	if err != nil {
		log.Fatal(err)
	}

	return *config
}
