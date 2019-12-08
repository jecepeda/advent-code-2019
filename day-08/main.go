package main

import (
	"fmt"
	"io/ioutil"
)

func readFile(filename string) ([][][]int, error) {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	result := make([][][]int, 0)
	for i := 0; i < len(raw)/150; i++ {
		instructions := make([][]int, 6)
		for idx := range instructions {
			instructions[idx] = make([]int, 25)
		}
		for j := 0; j < 6; j++ {
			for k := 0; k < 25; k++ {
				val := 0
				switch raw[(i*150)+(j*25)+k] {
				case '0':
					val = 0
				case '1':
					val = 1
				case '2':
					val = 2
				}
				if err != nil {
					return nil, err
				}
				instructions[j][k] = val
			}
		}
		result = append(result, instructions)
	}
	return result, nil
}

func getCount(layer [][]int, n int) int {
	result := 0
	for i := range layer {
		for j := range layer[i] {
			if layer[i][j] == n {
				result++
			}
		}
	}
	return result
}

func firstPart(picture [][][]int) {
	lowestCount := 0
	lowestIndex := -1
	for idx, layer := range picture {
		count := getCount(layer, 0)
		if count <= lowestCount || lowestIndex == -1 {
			lowestCount = count
			lowestIndex = idx
		}
	}
	ones := getCount(picture[lowestIndex], 1)
	twos := getCount(picture[lowestIndex], 2)
	fmt.Println("lowest index is", lowestIndex, lowestCount)
	fmt.Println("result is", ones*twos)
}

func secondPart(picture [][][]int) {
	finalImage := make([]byte, 150)
	for i := range finalImage {
		finalImage[i] = 2
	}
	for _, layer := range picture {
		for i := range layer {
			for j := range layer[i] {
				if finalImage[(i*25)+j] == 2 {
					finalImage[(i*25)+j] = byte(layer[i][j])
				}
			}
		}
	}
	for i := 0; i < 6; i++ {
		for j := 0; j < 25; j++ {
			fmt.Printf(" %v", finalImage[i*25+j])
		}
		fmt.Println()
	}
}

func main() {
	picture, err := readFile("input.txt")
	if err != nil {
		panic(err)
	}
	firstPart(picture)
	secondPart(picture)
}
