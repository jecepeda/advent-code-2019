package main

import "fmt"

func GetRange(low, high int) []int {
	result := make([]int, high-low+1)
	for i := 0; i < (high - low + 1); i++ {
		result[i] = low + i
	}
	return result
}
func ToString(intList []int) []string {
	result := make([]string, len(intList))
	for idx, i := range intList {
		result[idx] = fmt.Sprint(i)
	}
	return result
}

func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func matchesPassword(s string) bool {
	var (
		hasCheckedFirstElem bool
		lastChar            rune
		double              bool
	)
	for _, i := range s {
		if !hasCheckedFirstElem {
			hasCheckedFirstElem = true
		} else if i < lastChar {
			return false
		}
		if lastChar == i {
			double = true
		}
		lastChar = i
	}
	fmt.Println(s, double)
	return double
}

func matchesPasswordPartTwo(s string) bool {
	var (
		hasCheckedFirstElem bool
		lastChar            rune
		double              bool
	)
	for _, i := range s {
		if !hasCheckedFirstElem {
			hasCheckedFirstElem = true
		} else if i < lastChar {
			return false
		}
		if lastChar == i {
			double = true
		}
		lastChar = i
	}
	if !double {
		return false
	}
	hasCheckedFirstElem = false
	count := 0
	repeatGroup := false
	// now let's check the groups
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			count++
		} else {
			if count == 1 {
				repeatGroup = true
			}
			count = 0
		}
	}
	if count == 1 {
		repeatGroup = true
	}
	return repeatGroup
}

func main() {
	intRange := GetRange(357253, 892942)
	stringRange := ToString(intRange)

	result := Filter(stringRange, matchesPassword)
	fmt.Println(len(result))
	result = Filter(stringRange, matchesPasswordPartTwo)
	fmt.Println(len(result))
}
