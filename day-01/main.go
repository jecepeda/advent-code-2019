package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func readFile(filename string) ([]int, error) {
	var result []int
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	splitted := strings.Split(string(raw), "\n")
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

func calcFuel(x int) int {
	return (x / 3) - 2
}

func partOne(lines []int) int {
	var result int
	for _, l := range lines {
		result += calcFuel(l)
	}
	return result
}

func partTwo(lines []int) int {
	var result int
	for _, l := range lines {
		fuel := calcFuel(l)
		for fuel > 0 {
			result += fuel
			fuel = calcFuel(fuel)
		}
	}
	return result
}

func main() {
	lines, err := readFile("input.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("Part one:", partOne(lines))
	fmt.Println("Part two:", partTwo(lines))
}
