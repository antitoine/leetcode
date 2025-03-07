package main

import (
	"fmt"
	"math"
)

func closestPrimes(left int, right int) []int {
	if right < 2 {
		return []int{-1, -1}
	}
	isComposite := make([]bool, right+1)
	var primes []int
	for i := 2; i <= right; i++ {
		if !isComposite[i] {
			primes = append(primes, i)
			for j := i * i; j <= right; j += i {
				isComposite[j] = true
			}
		}
	}
	var filteredPrimes []int
	for _, p := range primes {
		if p >= left {
			filteredPrimes = append(filteredPrimes, p)
		}
	}
	if len(filteredPrimes) < 2 {
		return []int{-1, -1}
	}
	minDiff := math.MaxInt
	result := []int{-1, -1}
	for i := 0; i < len(filteredPrimes)-1; i++ {
		diff := filteredPrimes[i+1] - filteredPrimes[i]
		if diff < minDiff {
			minDiff = diff
			result[0] = filteredPrimes[i]
			result[1] = filteredPrimes[i+1]
		}
	}
	return result
}

func main() {
	left := 10
	right := 19
	result := closestPrimes(left, right)
	fmt.Printf("Result: %#v\n", result)
	if fmt.Sprintf("%v", result) == "[11 13]" {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
