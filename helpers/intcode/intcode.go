package intcode

import (
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

// InstructionLength maps an opcode to the number of parameters it has
var InstructionLength = map[int]int{
	1: 4,
	2: 4,
	3: 2,
	4: 2,
	5: 3,
	6: 3,
	7: 4,
	8: 4,
}

// Instructions is a bunch of intcode instructions
type Instructions []int

// Operation represents an OpCode along with the mdoes for its opcodes
type Operation struct {
	Modes  int
	OpCode int
}

// ReadInput from a file and return a set of Instructions
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

// Process the set of instructions
func (inst *Instructions) Process(input int) error {
	var opCodePos int

	for opCodePos = 0; (*inst)[opCodePos] != 99; {

		op := Operation{}
		op.Modes, op.OpCode = parseInstruction((*inst)[opCodePos])
		offset := opCodePos + InstructionLength[op.OpCode] - 1
		offsetVal := (*inst)[offset]
		jumped := false

		switch op.OpCode {
		case 1:
			params := inst.parseParams(op.Modes, opCodePos, (*inst)[opCodePos+1], (*inst)[opCodePos+2])
			(*inst)[offsetVal] = params[0] + params[1]
		case 2:
			params := inst.parseParams(op.Modes, opCodePos, (*inst)[opCodePos+1], (*inst)[opCodePos+2])
			(*inst)[offsetVal] = params[0] * params[1]
		case 3:
			(*inst)[offsetVal] = input
		case 4:
			params := inst.parseParams(op.Modes, opCodePos, (*inst)[opCodePos+1])
			logrus.Info(params[0])
		case 5:
			params := inst.parseParams(op.Modes, opCodePos, (*inst)[opCodePos+1], (*inst)[opCodePos+2])
			if params[0] != 0 {
				opCodePos = params[1]
				jumped = true
			}
		case 6:
			params := inst.parseParams(op.Modes, opCodePos, (*inst)[opCodePos+1], (*inst)[opCodePos+2])
			if params[0] == 0 {
				opCodePos = params[1]
				jumped = true
			}
		case 7:
			params := inst.parseParams(op.Modes, opCodePos, (*inst)[opCodePos+1], (*inst)[opCodePos+2])
			if params[0] < params[1] {
				(*inst)[offsetVal] = 1
			} else {
				(*inst)[offsetVal] = 0
			}
		case 8:
			params := inst.parseParams(op.Modes, opCodePos, (*inst)[opCodePos+1], (*inst)[opCodePos+2])
			if params[0] == params[1] {
				(*inst)[offsetVal] = 1
			} else {
				(*inst)[offsetVal] = 0
			}
		default:
			logrus.Fatal("Idk what to do with this opcode: ", op.OpCode)
		}

		if !jumped {
			opCodePos += InstructionLength[op.OpCode]
		}
	}

	return nil
}

func parseInstruction(inst int) (modes int, opCode int) {
	return inst / 100, inst % 100
}

// adapted from https://github.com/lizthegrey/adventofcode/blob/master/2019/intcode/vm.go#L59
func (inst *Instructions) parseParams(pModes int, opCodePos int, params ...int) (ints []int) {
	ints = make([]int, len(params))
	for i := 0; i < len(params); i++ {
		value := params[i]
		if pModes%10 == 0 {
			// position
			ints[i] = (*inst)[value]
		} else {
			// literal
			ints[i] = value
		}
		pModes /= 10
	}

	return ints
}
