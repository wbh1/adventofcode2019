package main

import (
	"github.com/sirupsen/logrus"
	"github.com/wbh1/adventofcode2019/helpers"
	"io/ioutil"
	"strconv"
)

var (
	fuelRequired = 0
)

// Integer division in Go will automatically round down
// Just account for
func calculateFuelRequired(mass int) int {
	req := (mass / 3) - 2
	if req >= 0 {
		return req
	}
	return 0
}

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		logrus.Fatal(err)
	}

	values := helpers.SplitOnNewLines(input)

	for _, v := range values {
		mass, _ := strconv.Atoi(v)
		requiredFuel := calculateFuelRequired(mass)
		fuelRequired += requiredFuel

		for requiredFuel > 0 {
			requiredFuel = calculateFuelRequired(requiredFuel)
			fuelRequired += requiredFuel
		}
	}

	logrus.Info("Fuel required: ", fuelRequired)
}
