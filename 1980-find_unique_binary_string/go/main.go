package main

import (
	"fmt"
)

func findDifferentBinaryString(nums []string) string {
	bitmask := 0
	for _, binaryString := range nums {
		bitmask |= 1 << countOnes(binaryString)
	}
	n := len(nums)
	for i := 0; i <= n; i++ {
		if bitmask>>i&1^1 == 1 {
			return generateBinaryString(i, n)
		}
	}
	return ""
}

func countOnes(s string) int {
	count := 0
	for _, char := range s {
		if char == '1' {
			count++
		}
	}
	return count
}

func generateBinaryString(ones, length int) string {
	result := ""
	for i := 0; i < ones; i++ {
		result += "1"
	}
	for i := ones; i < length; i++ {
		result += "0"
	}
	return result
}

func main() {
	nums := []string{"01", "10"}
	result := findDifferentBinaryString(nums)
	fmt.Printf("Result: %#v\n", result)
	if result == "00" {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
