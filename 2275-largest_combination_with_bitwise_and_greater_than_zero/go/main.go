package main

import (
	"fmt"
	"slices"
	"strconv"
)

func largestCombination(candidates []int) int {
	computedSetBits := make(map[int][]int)
	setBits := func(num int) []int {
		if bits, found := computedSetBits[num]; found {
			return bits
		}
		numBytes := strconv.FormatInt(int64(num), 2)
		var bits []int
		for i, b := range numBytes {
			if b == '1' {
				bits = append(bits, len(numBytes)-i-1)
			}
		}
		computedSetBits[num] = bits
		return bits
	}

	nbWithBitsAtIdx := make([]int, 24)
	for _, num := range candidates {
		bits := setBits(num)
		for _, idx := range bits {
			nbWithBitsAtIdx[idx]++
		}
	}

	return slices.Max(nbWithBitsAtIdx)
}

func main() {
	input := []int{16, 17, 71, 62, 12, 24, 14}
	result := largestCombination(input)
	fmt.Printf("Result: %#v\n", result)
	if result == 4 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
