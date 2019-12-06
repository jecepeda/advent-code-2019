package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

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
	Finish      Opcode = 99
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

// Execution is the function that executes a instruction
// at a given position
type Execution func(pos *int, i Instruction, instructions []int) ([]int, bool)

var executions = map[Opcode]Execution{
	Sum:         execSum,
	Multiply:    execMultiplication,
	Input:       execInput,
	Output:      execOutput,
	Finish:      execBreak,
	JumpIfTrue:  execJumpIfTrue,
	JumpIfFalse: execJumpIfFalse,
	LessThan:    execLessThan,
	Equals:      execEquals,
}

func readInstruction(ins int) Instruction {
	return Instruction{
		Type:   Opcode(ins % 100),
		First:  ParameterMode((ins % 1000) / 100),
		Second: ParameterMode((ins % 10000) / 1000),
		Output: ParameterMode((ins % 100000) / 10000),
	}
}

func readFile(filename string) ([]int, error) {
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
	case Finish:
		return 0
	}
	return 0
}

func execSum(pos *int, i Instruction, instructions []int) ([]int, bool) {
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

func execMultiplication(pos *int, i Instruction, instructions []int) ([]int, bool) {
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

func execInput(pos *int, i Instruction, instructions []int) ([]int, bool) {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter code: ")
	reader.Scan()
	text := reader.Text()
	val, _ := strconv.Atoi(text)
	i1 := instructions[*pos+1]
	instructions[i1] = val
	*pos += increaseCounter(i)
	return instructions, false
}

func execOutput(pos *int, i Instruction, instructions []int) ([]int, bool) {
	i1 := instructions[*pos+1]
	if i.First == PositionMode {
		i1 = instructions[i1]
	}
	fmt.Println("output", i1)
	*pos += increaseCounter(i)
	return instructions, false
}

func execBreak(pos *int, i Instruction, instructions []int) ([]int, bool) {
	return instructions, true
}

func execJumpIfTrue(pos *int, i Instruction, instructions []int) ([]int, bool) {
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

func execJumpIfFalse(pos *int, i Instruction, instructions []int) ([]int, bool) {
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

func execLessThan(pos *int, i Instruction, instructions []int) ([]int, bool) {
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

func execEquals(pos *int, i Instruction, instructions []int) ([]int, bool) {
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

func execOpCodes(instructions []int) []int {
	i := 0
	var shouldBreak bool
	for i < len(instructions) {
		instruction := readInstruction(instructions[i])
		instructions, shouldBreak = executions[instruction.Type](&i, instruction, instructions)
		if shouldBreak {
			break
		}
	}
	return instructions
}

func main() {
	instructions, err := readFile("input.txt")
	if err != nil {
		panic(err)
	}
	execOpCodes(instructions)
}
