package main

import (
	"fmt"
	"sort"
)

type Pair struct {
	Value int
	Index int
}

func lexicographicallySmallestArray(nums []int, limit int) []int {
	length := len(nums)
	numWithIndex := make([]Pair, length)
	for i, num := range nums {
		numWithIndex[i] = Pair{Value: num, Index: i}
	}
	sort.Slice(numWithIndex, func(i, j int) bool {
		return numWithIndex[i].Value < numWithIndex[j].Value
	})
	result := make([]int, length)
	i := 0
	for i < length {
		j := i + 1
		for j < length && numWithIndex[j].Value-numWithIndex[j-1].Value <= limit {
			j++
		}
		indices := make([]int, j-i)
		for k := i; k < j; k++ {
			indices[k-i] = numWithIndex[k].Index
		}
		sort.Ints(indices)
		for k, idx := range indices {
			result[idx] = numWithIndex[i+k].Value
		}
		i = j
	}

	return result
}

func main() {
	nums := []int{1, 5, 3, 9, 8}
	limit := 2
	result := lexicographicallySmallestArray(nums, limit)
	fmt.Printf("Result: %#v\n", result)
	if fmt.Sprintf("%v", result) == "[1 3 5 8 9]" {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
