package main

import (
	"github.com/sirupsen/logrus"
	"strconv"
)

const (
	min = 165432
	max = 707912
)

var partA, partB int

func main() {
	for i := min; i < max; i++ {
		if !decreases(i) && containsTwoAdjacentDigits(i, 'A') {
			partA++
		}
		if !decreases(i) && containsTwoAdjacentDigits(i, 'B') {
			partB++
		}
	}

	logrus.Infof("Part A: %d; Part B: %d", partA, partB)
}

func decreases(n int) bool {
	nStr := strconv.Itoa(n)
	prev := int32(byte(0))
	for _, v := range nStr {
		if v < prev {
			return true
		}
		prev = v
	}
	return false
}

func containsTwoAdjacentDigits(n int, part rune) bool {
	digits := make(map[int32]int)
	nStr := strconv.Itoa(n)

	for _, v := range nStr {
		digits[v] += 1
	}

	for _, occurrences := range digits {
		if occurrences >= 2 && part == 'A' {
			return true
		} else if occurrences == 2 && part == 'B' {
			return true
		}
	}

	return false
}
