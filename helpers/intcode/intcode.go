package intcode

import (
	"strconv"
	"strings"
)
var OpCodeParams = map[int]int{
	1:  3,
	2: 3,
	3: 1,
	4: 1,
}
type Instructions []int

func ReadInput(program []byte) (Instructions, error) {
	var instructions Instructions
	inst := strings.Split(string(program), ",")

	for _, v := range inst {
		i, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		instructions = append(instructions, i)
	}

	return instructions, nil
}

func (inst *Instructions) Process() error {
	var opCodePos int

	for _, i := range *inst {
		if i == 99 {
			return nil
		}

		opCode =
	}
}