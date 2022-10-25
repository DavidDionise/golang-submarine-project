package main

import (
	"com.pi/submarine/compass"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	c := compass.Compass{}
	c.Start()
}
