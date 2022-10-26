package steering

import (
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestCalculateAngleOfDeviation(t *testing.T) {

	// 90 degrees to the right - Q1
	var currentRead = 0
	var setHead = 90

	var expectedValue = -90
	var actualValue = calculateAngleOfDeviation(currentRead, setHead)

	if expectedValue != actualValue {
		log.Fatal(expectedValue, " does not equal ", actualValue)
	}

	// 10 degrees to the left - Q1
	currentRead = 90
	setHead = 80

	expectedValue = 10
	actualValue = calculateAngleOfDeviation(currentRead, setHead)

	if expectedValue != actualValue {
		log.Fatal(expectedValue, " does not equal ", actualValue)
	}

	// 20 degrees to the left - Q1 -> Q2
	currentRead = 10
	setHead = 350

	expectedValue = 20
	actualValue = calculateAngleOfDeviation(currentRead, setHead)

	if expectedValue != actualValue {
		log.Fatal(expectedValue, " does not equal ", actualValue)
	}

	// 50 degrees to the right - Q2 -> Q1
	currentRead = 340
	setHead = 30

	expectedValue = -50
	actualValue = calculateAngleOfDeviation(currentRead, setHead)

	if expectedValue != actualValue {
		log.Fatal(expectedValue, " does not equal ", actualValue)
	}
}
