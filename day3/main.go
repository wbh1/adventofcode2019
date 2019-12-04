package main

import (
	"github.com/sirupsen/logrus"
	"github.com/wbh1/adventofcode2019/helpers"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

var (
	line1 []string
	line2 []string
	grid  [10001][10001]int
	xPos  = 5001
	yPos  = 5001
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		logrus.Error(err)
	}

	directions := helpers.SplitOnNewLines(input)

	line1 = make([]string, len(directions[0]))
	line2 = make([]string, len(directions[1]))

	for i, v := range strings.Split(string(directions[0]), ",") {
		line1[i] = v
	}

	for i, v := range strings.Split(string(directions[1]), ",") {
		line2[i] = v
	}

	part1()
}

func part1() {
	for _, v := range line1 {
		if v == "" {
			continue
		}
		if xAxis(v) {
			xPos = xPos + movement(v)
			grid[xPos][yPos] = 1
		} else {
			yPos = yPos + movement(v)
			grid[xPos][yPos+movement(v)] = 1
		}
	}

	for _, v := range line2 {
		if v == "" {
			continue
		}
		if xAxis(v) {
			xPos = xPos + movement(v)
			if grid[xPos][yPos] == 1 {
				grid[xPos][yPos] += 2
			}
		} else {
			yPos = yPos + movement(v)
			if grid[xPos][yPos] == 1 {
				grid[xPos][yPos] += 2
			}
		}
	}

	for y, row := range grid[:] {
		for x, v := range row {
			if v == 3 {
				logrus.Info(x, y,
					math.Abs(float64(x)-5001.0)+math.Abs(float64(y)-5001.0))
			}
		}
	}
}

func xAxis(direction string) bool {
	if strings.Contains(direction, "L") {
		return true
	}
	if strings.Contains(direction, "R") {
		return true
	}
	return false
}

func movement(direction string) int {
	logrus.Info(direction)
	spacesMoved := []rune(direction)[1:]
	move, err := strconv.Atoi(string(spacesMoved))
	if err != nil {
		logrus.Error(err)
	}

	if strings.Contains(direction, "L") {
		return 0 - move
	}
	if strings.Contains(direction, "R") {
		return move
	}

	return 0
}
