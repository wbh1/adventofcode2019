package main

import (
	"github.com/sirupsen/logrus"
	"github.com/wbh1/adventofcode2019/helpers/intcode"
	"io/ioutil"
	"strconv"
	"strings"
)

var (
	opcodePos int
	opcode int
	input1 int
	input2 int
	output int
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		logrus.Fatal(err)
	}

	strNums := strings.Split(string(input), ",")
	numbers := make([]int, len(strNums))
	for i, v := range strNums {
		num, _ := strconv.Atoi(v)
		numbers[i] = num
	}

	numbers = part1(numbers)
	numbers1, err := intcode.ReadInput(input)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Info(numbers1)
	numbers1.Process()
	logrus.Info(numbers[0], numbers1[0])

	//param1, param2 := part2(numbers)
	//logrus.Info(param1, param2)
}

func compute(opcode, num1, num2 int) int {
	if opcode == 1 {
		logrus.Infof("Adding %d and %d", num1, num2)
		return num1 + num2
	} else if opcode == 2 {
		logrus.Infof("Multiplying %d and %d", num1, num2)
		return num1 * num2
	} else {
		logrus.Fatal(opcode, " is not a valid opcode")
	}
	return 0
}

func part1(numbers []int) []int {
	opcode = numbers[opcodePos]
	for opcode != 99 {
		input1 = numbers[numbers[opcodePos+1]]
		input2 = numbers[numbers[opcodePos+2]]
		numbers[numbers[opcodePos+3]] = compute(opcode, input1, input2)
		opcodePos+=4
		opcode = numbers[opcodePos]
	}

	return numbers
}

func part2(numbers []int) (param1, param2 int) {
	output := 0
	for param1 := 0; param1 <= 99 && output != 19690720; param1++ {
		for param2 := 0; param2 <= 99 && output != 19690720; param2++ {
			opcodePos = 0
			numbers = readInput()
			numbers[1] = param1
			numbers[2] = param2
			output = part1(numbers)[0]
		}
	}
	return numbers[1], numbers[2]
}

func readInput() []int {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		logrus.Fatal(err)
	}

	strNums := strings.Split(string(input), ",")
	numbers := make([]int, len(strNums))
	for i, v := range strNums {
		num, _ := strconv.Atoi(v)
		numbers[i] = num
	}

	return numbers
}