package compass

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/tkanos/gonfig"
	"periph.io/x/conn/v3/i2c"
	"periph.io/x/conn/v3/i2c/i2creg"
	"periph.io/x/host/v3"
)

type Compass struct {
	angle int
}

func (c Compass) Start() {

	config := CompassConfig{}
	config.start()

	fmt.Println("address: ", config.I2cAddress)
	fmt.Println("milli: ", config.FetchRateMiliseconds)

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

		time.Sleep(time.Duration(config.FetchRateMiliseconds * 1000000))
	}
}

type CompassConfig struct {
	I2cAddress           string
	FetchRateMiliseconds int
}

func (c *CompassConfig) start() {
	dir, err := os.Getwd()

	if err != nil {
		log.Fatalln(err)
	}

	err = gonfig.GetConf(dir+"/compass/config.json", c)

	if err != nil {
		log.Fatal(err)
	}
}
