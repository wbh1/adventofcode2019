package main

import (
	"github.com/sirupsen/logrus"
	"github.com/wbh1/adventofcode2019/helpers"
	"io/ioutil"
	"strings"
)

var (
	orbits = make(map[string]string)
	distanceToCenter = make(map[string]int)
	allDistances int
)

func main()  {

	f, err := ioutil.ReadFile("input.txt")
	if err != nil {
		logrus.Fatal(err)
	}

	// Create a map of orbits in which the key is the planet orbiting
	// and the value is the planet being orbited
	for _, v := range helpers.SplitOnNewLines(f) {
		objects := strings.Split(v, ")")
		center := objects[0]
		orbiter := objects[1]
		orbits[orbiter] = center
	}

	// for each object in an orbit, calculate its distance to the COM
	// and add it to the allDistances var
	distanceToCenter["COM"] = 0
	for k := range orbits {
		allDistances += findCOMDistance(k)
	}

	logrus.Infof("Part1: %d", allDistances)

	// Determine what I orbit (directly/indirectly) and what Santa orbits
	myOrbits := make([]string, 0)
	santaOrbits := make([]string, 0)
	for loc := "YOU"; loc != "COM"; loc = orbits[loc] {
		myOrbits = append(myOrbits, loc)
	}
	for loc := "SAN"; loc != "COM"; loc = orbits[loc] {
		santaOrbits = append(santaOrbits, loc)
	}

	// Work backwards through our orbits until we find our common point.
	// Then, `i` will be equal to the length of our orbits and the length of Santa's
	// minus the number of steps times 2 (since we are stepping back 1 in *both* Santa's orbit and ours).
	for i := 1; ; i++ {
		me := myOrbits[len(myOrbits)-i]
		santa := santaOrbits[len(santaOrbits)-i]
		if me != santa {
			logrus.Infof("Part B: %d", len(myOrbits)+len(santaOrbits)-(2*i))
			break
		}
	}
}

func findCOMDistance(object string) int {
	if _, ok := distanceToCenter[object]; !ok {
		distanceToCenter[object] = findCOMDistance(orbits[object]) + 1
	}
	return distanceToCenter[object]
}