package main

import (
	"github.com/sirupsen/logrus"
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

	opcode = numbers[opcodePos]
	for opcode != 99 {
		input1 = numbers[numbers[opcodePos+1]]
		input2 = numbers[numbers[opcodePos+2]]
		numbers[numbers[opcodePos+3]] = compute(opcode, input1, input2)
		opcodePos+=4
		opcode = numbers[opcodePos]

	}

	logrus.Info(numbers)
}

func compute(opcode, num1, num2 int) int {
	if opcode == 1 {
		return num1 + num2
	} else if opcode == 2 {
		return num1 * num2
	} else {
		logrus.Fatal(opcode, " is not a valid opcode")
	}
	return 0
}
