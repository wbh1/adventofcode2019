package main

import (
	"github.com/sirupsen/logrus"
	"github.com/wbh1/adventofcode2019/helpers/intcode"
	"io/ioutil"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		logrus.Fatal(err)
	}

	inst, err := intcode.ReadInput(input)
	if err != nil {
		logrus.Fatal("Couldn't process intcode ", err)
	}
	logrus.Info(inst)
}
