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
	x              int
	y              int
	distFromOrigin float64
}

var (
	updatePoint func(p *point)
)

func main() {

	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		logrus.Error("Error reading file: ", err)
	}

	contents := helpers.SplitOnNewLines(bytes)
	lines := make(map[point][2]float64)

	for i, line := range contents {
		curPoint := point{0, 0, 0}
		traveled := 0.0
		for _, step := range strings.Split(line, ",") {
			dist, err := strconv.Atoi(step[1:])
			if err != nil {
				logrus.Fatal("Couldn't parse step: ", step)
			}

			switch step[0] {
			case 'L':
				updatePoint = func(p *point) {
					p.x--
				}
			case 'R':
				updatePoint = func(p *point) {
					p.x++
				}
			case 'U':
				updatePoint = func(p *point) {
					p.y++
				}
			case 'D':
				updatePoint = func(p *point) {
					p.y--
				}
			default:
				logrus.Fatal("Uh oh... couldn't parse this step: ", step)
			}

			for mv := 1; mv <= dist; mv++ {
				updatePoint(&curPoint)
				hits := lines[curPoint]
				if hits[i] != 0 {
					hits[i] = math.Min(hits[i], float64(mv)+traveled)
					logrus.Info(hits)
				} else {
					hits[i] = float64(mv) + traveled
				}
				lines[curPoint] = hits
			}
			traveled += float64(dist)

		}
	}
	closest := math.MaxFloat64
	least := math.MaxFloat64
	for loc, hits := range lines {
		if hits[0] != float64(0) && hits[1] != float64(0) {
			distance := math.Abs(float64(loc.x)) + math.Abs(float64(loc.y))
			if distance < closest {
				closest = distance
			}
			sum := hits[0] + hits[1]
			if sum < least {
				least = sum
			}
		}
	}

	logrus.Infof("Part A: %d; Part B: %d", int(closest), int(least))
}