package main

import (
	"fmt"

	"gonum.org/v1/gonum/stat/combin"

	"github.com/jecepeda/advent-code-2019/intcode"
)

func copyInstructions(instructions []int) []int {
	result := make([]int, len(instructions))
	copy(result, instructions)
	return result
}

func getPermutations(orig []int, perms [][]int) [][]int {
	result := make([][]int, len(perms))
	for idx, p := range perms {
		result[idx] = make([]int, len(p))
		copy(result[idx], p)
		for i := 0; i < len(p); i++ {
			result[idx][i] = orig[p[i]]
		}
	}
	return result
}

func partOne(instructions []int) {
	amplifiers := make([]*intcode.Intcode, 5)
	for i := range amplifiers {
		amplifiers[i] = intcode.NewIntCodeProgram()
	}

	var firstOutput int
	firstPermutations := combin.Permutations(5, 5)
	for _, comb := range firstPermutations {
		phase := 0
		for i := 0; i < 5; i++ {
			amplifiers[i].Reset()
			amplifiers[i].Input <- comb[i]
			amplifiers[i].Input <- phase
			if !amplifiers[i].Started {
				go amplifiers[i].Exec(copyInstructions(instructions))
			}
			phase = <-amplifiers[i].Output
			_ = amplifiers[i].Finish
		}
		output := amplifiers[len(amplifiers)-1].Out
		if output > firstOutput {
			firstOutput = output
		}
	}
	fmt.Println("first output is", firstOutput)
}

func partTwo(instructions []int) {
	amplifiers := make([]*intcode.Intcode, 5)
	var secondOutput int
	permutations := combin.Permutations(5, 5)
	transformed := getPermutations([]int{5, 6, 7, 8, 9}, permutations)
	for _, comb := range transformed {
		phase := 0
		finish := false
		for i := range amplifiers {
			amplifiers[i] = intcode.NewIntCodeProgram()
		}
		for !finish {
			for i := 0; i < 5; i++ {
				select {
				case _ = <-amplifiers[i].Finish:
					if i == 4 {
						finish = true
						break
					}
				default:
					if !amplifiers[i].Started {
						amplifiers[i].Input <- comb[i]
						go amplifiers[i].Exec(copyInstructions(instructions))
					}
					amplifiers[i].Input <- phase
					phase = <-amplifiers[i].Output
				}
			}
		}
		out := amplifiers[len(amplifiers)-1].Out
		if out > secondOutput {
			secondOutput = out
		}
	}
	fmt.Println("second output is", secondOutput)
}

func main() {
	instructions, err := intcode.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	partOne(instructions)
	partTwo(instructions)
}
