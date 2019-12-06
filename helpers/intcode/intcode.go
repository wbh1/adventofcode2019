package intcode

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
)
var InstructionLength = map[int]int{
	1:  4,
	2: 4,
	3: 2,
	4: 2,
}

type Instructions []int

type Operation struct {
	Modes ParamModes
	OpCode int
}

type ParamModes struct {
	Param1 int
	Param2 int
	Param3 int
}

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

	input := 1


	for opCodePos = 0; (*inst)[opCodePos] != 99; {

		op := Operation{}

		op.Modes, op.OpCode = parseInstruction((*inst)[opCodePos])
		logrus.Infof("Opcode: %d; instruction: %d; modes: %v", op.OpCode, (*inst)[opCodePos], op.Modes)

		switch op.OpCode {
		case 1:
			params := inst.parseParams(op.Modes, (*inst)[opCodePos+1], (*inst)[opCodePos+2])
			bah := *inst
			logrus.Infof("Adding %d and %d; storing at %d in place of %d. Modes: %v", params[0], params[1], bah[opCodePos+3], bah[bah[opCodePos+3]], op.Modes)
			(*inst)[(*inst)[opCodePos+3]] = params[0] + params[1]
			opCodePos += InstructionLength[1]
		case 2:
			params := inst.parseParams(op.Modes, (*inst)[opCodePos+1], (*inst)[opCodePos+2])
			logrus.Infof("Multiplying %d and %d; storing at %d in place of %d. Modes: %v", params[0], params[1], (*inst)[opCodePos+3], (*inst)[(*inst)[opCodePos+3]], op.Modes)
			(*inst)[(*inst)[opCodePos+3]] = params[0] * params[1]
			opCodePos += InstructionLength[2]
		case 3:
			logrus.Infof("Storing %d at pos %d", input, (*inst)[opCodePos+1])
			(*inst)[(*inst)[opCodePos+1]] = input
			opCodePos += InstructionLength[3]
		case 4:
			logrus.Info((*inst)[(*inst)[opCodePos+1]])
			opCodePos += InstructionLength[4]
		default:
			logrus.Fatal("Idk what to do with this opcode: ", op.OpCode)
		}
	}

	return nil
}

func parseInstruction(inst int) (modes ParamModes, opCode int) {
	str := strconv.Itoa(inst)
	opCode, err := strconv.Atoi(str[len(str)-1:])
	if err != nil {
		logrus.Fatal(err)
	}

	str = fmt.Sprintf("|%05s|", str)
	for i := len(str)-3; i >= 0; i-- {
		//logrus.Infof("i is %d. inst is %v", i, inst)
		switch i {
		case 3:
			modes.Param1 = mode(str[i])
		case 2:
			modes.Param2 = mode(str[i])
		case 1:
			modes.Param3 = mode(str[i])
		}
	}

	return
}

func mode(num uint8) int {
	switch num {
	case 48:
		return 0
	case 49:
		return 1
	default:
		logrus.Fatal("unsupported mode: ", num)
	}
	return 39
}

func (inst *Instructions) parseParams(modes ParamModes, params ...int) (ints []int) {
	logrus.Info("params: ", params)
	for i, p := range params {
		switch i {
		case 0:
			if modes.Param1 == 1 {
				ints = append(ints, p)
			} else {
				ints = append(ints, (*inst)[p])
			}
		case 1:
			if modes.Param2 == 1 {
				ints = append(ints, p)
			} else {
				ints = append(ints, (*inst)[p])
			}
		case 2:
			if modes.Param3 == 1 {
				ints = append(ints, p)
			} else {
				ints = append(ints, (*inst)[p])
			}
		}
	}

	return
}