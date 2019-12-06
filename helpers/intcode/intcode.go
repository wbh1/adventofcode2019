package intcode

import (
	"github.com/sirupsen/logrus"
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

	for {
		var (
			op Operation
			i = (*inst)[opCodePos]
		)

		if i == 99 {
			return nil
		}

		//logrus.Info(*inst)
		op.Modes, op.OpCode = parseInstruction(i)
		logrus.Info(op.Modes)

		switch op.OpCode {
		case 1:
			params := inst.parseParams(op.Modes, (*inst)[opCodePos+1], (*inst)[opCodePos+2])
			//logrus.Infof("Adding %d and %d; storing at %d in place of %d", params[0], params[1], opCodePos+3, (*inst)[opCodePos+3])
			(*inst)[(*inst)[opCodePos+3]] = params[0] + params[1]
			opCodePos += 4
		case 2:
			params := inst.parseParams(op.Modes, (*inst)[opCodePos+1], (*inst)[opCodePos+2])
			//logrus.Infof("Multiplying %d and %d; storing at %d in place of %d", params[0], params[1], opCodePos+3, (*inst)[opCodePos+3])
			(*inst)[(*inst)[opCodePos+3]] = params[0] * params[1]
			opCodePos += 4
		case 3:
			//logrus.Infof("Storing %d at pos %d", input, (*inst)[opCodePos+1])
			(*inst)[(*inst)[opCodePos+1]] = input
			opCodePos += 2
		case 4:
			params := inst.parseParams(op.Modes, (*inst)[opCodePos+1])
			logrus.Info((*inst)[params[0]])
			opCodePos += 2
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

	for i := len(str)-2; i >= 0; i-- {
		switch i {
		case 3:
			logrus.Infof("Param3 = %d", int(str[0]))
			modes.Param3 = int(str[0])
		case 2:
			logrus.Infof("Param2 = %d", int(str[1]))
			modes.Param2 = int(str[1])
		case 1:
			logrus.Infof("Param1 = %d", int(str[2]))
			modes.Param1 = mode(str[2])
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
	for i, p := range params {
		switch i {
		case 0:
			if modes.Param1 == 1 {
				ints = append(ints, (*inst)[p])
			} else {
				ints = append(ints, p)
			}
		case 1:
			if modes.Param2 == 1 {
				ints = append(ints, (*inst)[p])
			} else {
				ints = append(ints, p)
			}
		case 2:
			if modes.Param3 == 1 {
				ints = append(ints, (*inst)[p])
			} else {
				ints = append(ints, p)
			}
		}
	}

	return
}