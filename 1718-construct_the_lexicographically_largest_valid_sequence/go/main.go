package main

import (
	"fmt"
)

func constructDistancedSequence(n int) []int {
	sequenceLength := n
	counter := make([]int, n*2)
	resultSequence := make([]int, n*2)
	for i := range counter {
		counter[i] = 2
	}
	counter[1] = 1

	var dfs func(position int) bool
	dfs = func(position int) bool {
		if position == sequenceLength*2 {
			return true
		}
		if resultSequence[position] != 0 {
			return dfs(position + 1)
		}
		for i := sequenceLength; i > 1; i-- {
			if counter[i] > 0 && position+i < sequenceLength*2 && resultSequence[position+i] == 0 {
				resultSequence[position] = i
				resultSequence[position+i] = i
				counter[i] = 0
				if dfs(position + 1) {
					return true
				}
				counter[i] = 2
				resultSequence[position] = 0
				resultSequence[position+i] = 0
			}
		}
		if counter[1] > 0 {
			resultSequence[position] = 1
			counter[1] = 0
			if dfs(position + 1) {
				return true
			}
			counter[1] = 1
			resultSequence[position] = 0
		}
		return false
	}

	dfs(1)
	ans := []int{}
	for i := 1; i < n*2; i++ {
		ans = append(ans, resultSequence[i])
	}
	return ans
}

func main() {
	n := 3
	result := constructDistancedSequence(n)
	fmt.Printf("Result: %#v\n", result)
	if fmt.Sprintf("%v", result) == "[3 1 2 3 2]" {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
