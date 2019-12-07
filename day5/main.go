package main

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
	"github.com/wbh1/adventofcode2019/helpers/intcode"
)

func main() {
	run(1, "input.txt")
	run(5, "input2.txt")

}

func run(input int, file string) {
	f, err := ioutil.ReadFile(file)
	if err != nil {
		logrus.Fatal(err)
	}

	inst, err := intcode.ReadInput(f)
	if err != nil {
		logrus.Fatal("Couldn't process intcode ", err)
	}

	if err := inst.Process(input); err != nil {
		logrus.Fatal(err)
	}
}
