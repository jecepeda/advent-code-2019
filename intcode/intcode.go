package intcode

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// Intcode is the main program that executes a set of instructions
type Intcode struct {
	Input    chan int
	Output   chan int
	Finish   chan bool
	Started  bool
	Finished bool
	Out      int
	name     int
	logger   func(val interface{})
}

// NewIntCodeProgram generates a new program
// with default input function initialized
func NewIntCodeProgram() *Intcode {
	icode := &Intcode{}
	icode.Reset()
	return icode
}

func (icode *Intcode) Reset() {
	in := make(chan int, 20)
	out := make(chan int)
	finish := make(chan bool)
	icode.Input = in
	icode.Output = out
	icode.Finish = finish
	icode.Started = false
	icode.Finished = false
}

func (icode *Intcode) Halt() {
	close(icode.Input)
	close(icode.Output)
	icode.Finish <- true
	icode.Finished = true
	icode.Started = true
}

// Opcode is the kinds of operations we can encounter in
// the instructions
type Opcode int

// possible operations
const (
	Sum         Opcode = 1
	Multiply    Opcode = 2
	Input       Opcode = 3
	Output      Opcode = 4
	JumpIfTrue  Opcode = 5
	JumpIfFalse Opcode = 6
	LessThan    Opcode = 7
	Equals      Opcode = 8
	Break       Opcode = 99
)

// ParameterMode is the kind of parameters we can
// encounter when reading operations
type ParameterMode int

// possible parameter modes
const (
	PositionMode  ParameterMode = 0
	InmediateMode ParameterMode = 1
)

// Instruction is the complete instruction
// we need to execute, with al parameters inside
type Instruction struct {
	Type   Opcode
	First  ParameterMode
	Second ParameterMode
	Output ParameterMode
}

func (i Instruction) String() string {
	return fmt.Sprintf(
		"Istruction: %d, %d, %d, %d",
		i.Type, i.First, i.Second, i.Output)
}

func readInstruction(ins int) Instruction {
	return Instruction{
		Type:   Opcode(ins % 100),
		First:  ParameterMode((ins % 1000) / 100),
		Second: ParameterMode((ins % 10000) / 1000),
		Output: ParameterMode((ins % 100000) / 10000),
	}
}

// ReadFile reads the file given by the filename
// and returns the possible instuctions to be executed
func ReadFile(filename string) ([]int, error) {
	var result []int
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	splitted := strings.Split(string(raw), ",")
	for _, s := range splitted {
		if s == "" {
			continue
		}
		v, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		result = append(result, v)
	}
	return result, nil
}

func increaseCounter(ins Instruction) int {
	switch ins.Type {
	case Sum, Multiply, LessThan, Equals:
		return 4
	case Input, Output:
		return 2
	case JumpIfTrue, JumpIfFalse:
		return 3
	case Break:
		return 0
	}
	return 0
}

func (icode *Intcode) execSum(pos *int, i Instruction, instructions []int) ([]int, bool) {
	i1, i2 := instructions[*pos+1], instructions[*pos+2]
	resultAddres := instructions[*pos+3]
	if i.First == PositionMode {
		i1 = instructions[i1]
	}
	if i.Second == PositionMode {
		i2 = instructions[i2]
	}
	instructions[resultAddres] = i1 + i2
	*pos += increaseCounter(i)
	return instructions, false
}

func (icode *Intcode) execMultiplication(pos *int, i Instruction, instructions []int) ([]int, bool) {
	i1, i2 := instructions[*pos+1], instructions[*pos+2]
	resultAddres := instructions[*pos+3]
	if i.First == PositionMode {
		i1 = instructions[i1]
	}
	if i.Second == PositionMode {
		i2 = instructions[i2]
	}
	instructions[resultAddres] = i1 * i2
	*pos += increaseCounter(i)
	return instructions, false
}

func (icode *Intcode) execInput(pos *int, i Instruction, instructions []int) ([]int, bool) {
	val := <-icode.Input
	i1 := instructions[*pos+1]
	instructions[i1] = val
	*pos += increaseCounter(i)
	return instructions, false
}

func (icode *Intcode) execOutput(pos *int, i Instruction, instructions []int) ([]int, bool) {
	i1 := instructions[*pos+1]
	if i.First == PositionMode {
		i1 = instructions[i1]
	}
	icode.Output <- i1
	icode.Out = i1
	*pos += increaseCounter(i)
	return instructions, false
}

func (icode *Intcode) execBreak(pos *int, i Instruction, instructions []int) ([]int, bool) {
	return instructions, true
}

func (icode *Intcode) execJumpIfTrue(pos *int, i Instruction, instructions []int) ([]int, bool) {
	i1, i2 := instructions[*pos+1], instructions[*pos+2]
	if i.First == PositionMode {
		i1 = instructions[i1]
	}
	if i.Second == PositionMode {
		i2 = instructions[i2]
	}
	if i1 != 0 {
		*pos = i2
	} else {
		*pos += increaseCounter(i)
	}
	return instructions, false
}

func (icode *Intcode) execJumpIfFalse(pos *int, i Instruction, instructions []int) ([]int, bool) {
	i1, i2 := instructions[*pos+1], instructions[*pos+2]
	if i.First == PositionMode {
		i1 = instructions[i1]
	}
	if i.Second == PositionMode {
		i2 = instructions[i2]
	}
	if i1 == 0 {
		*pos = i2
	} else {
		*pos += increaseCounter(i)
	}
	return instructions, false
}

func (icode *Intcode) execLessThan(pos *int, i Instruction, instructions []int) ([]int, bool) {
	i1, i2 := instructions[*pos+1], instructions[*pos+2]
	resultAddres := instructions[*pos+3]
	if i.First == PositionMode {
		i1 = instructions[i1]
	}
	if i.Second == PositionMode {
		i2 = instructions[i2]
	}
	var val int
	if i1 < i2 {
		val = 1
	}
	instructions[resultAddres] = val
	*pos += increaseCounter(i)
	return instructions, false
}

func (icode *Intcode) execEquals(pos *int, i Instruction, instructions []int) ([]int, bool) {
	i1, i2 := instructions[*pos+1], instructions[*pos+2]
	resultAddres := instructions[*pos+3]
	if i.First == PositionMode {
		i1 = instructions[i1]
	}
	if i.Second == PositionMode {
		i2 = instructions[i2]
	}
	var val int
	if i1 == i2 {
		val = 1
	}
	instructions[resultAddres] = val
	*pos += increaseCounter(i)
	return instructions, false
}

// Exec execs the instructions passed by argument
func (icode *Intcode) Exec(instructions []int) int {
	icode.Started = true
	i := 0
	var shouldBreak bool
	for i < len(instructions) {
		instruction := readInstruction(instructions[i])
		switch instruction.Type {
		case Sum:
			instructions, shouldBreak = icode.execSum(&i, instruction, instructions)
		case Multiply:
			instructions, shouldBreak = icode.execMultiplication(&i, instruction, instructions)
		case Input:
			instructions, shouldBreak = icode.execInput(&i, instruction, instructions)
		case Output:
			instructions, shouldBreak = icode.execOutput(&i, instruction, instructions)
		case JumpIfTrue:
			instructions, shouldBreak = icode.execJumpIfTrue(&i, instruction, instructions)
		case JumpIfFalse:
			instructions, shouldBreak = icode.execJumpIfFalse(&i, instruction, instructions)
		case LessThan:
			instructions, shouldBreak = icode.execLessThan(&i, instruction, instructions)
		case Equals:
			instructions, shouldBreak = icode.execEquals(&i, instruction, instructions)
		case Break:
			instructions, shouldBreak = icode.execBreak(&i, instruction, instructions)
		}
		if shouldBreak {
			break
		}
	}
	icode.Halt()
	return icode.Out
}
