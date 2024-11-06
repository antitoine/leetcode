package main

import (
	"fmt"
	"sort"
	"strconv"
)

func canSortArray(nums []int) bool {
	computedIntSetBits := make(map[int]int)
	nbSetBits := func(num int) int {
		if nb, found := computedIntSetBits[num]; found {
			return nb
		}
		numBytes := strconv.FormatInt(int64(num), 2)
		nb := 0
		for _, b := range numBytes {
			if b == '1' {
				nb++
			}
		}
		computedIntSetBits[num] = nb
		return nb
	}

	k := 0
	seqences := [][]int{{nums[0]}}
	for i := 1; i < len(nums); i++ {
		previousSequenceSetBits := nbSetBits(seqences[k][0])
		numSetBits := nbSetBits(nums[i])
		if previousSequenceSetBits == numSetBits {
			seqences[k] = append(seqences[k], nums[i])
		} else {
			seqences = append(seqences, []int{nums[i]})
			k++
			sort.Ints(seqences[k-1])
			if k > 1 && seqences[k-2][len(seqences[k-2])-1] > seqences[k-1][0] {
				return false
			}
		}
	}
	sort.Ints(seqences[k])
	if k > 0 && seqences[k-1][len(seqences[k-1])-1] > seqences[k][0] {
		return false
	}
	return true
}

func main() {
	input := []int{8, 4, 2, 30, 15}
	result := canSortArray(input)
	fmt.Printf("Result: %#v\n", result)
	if result {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
