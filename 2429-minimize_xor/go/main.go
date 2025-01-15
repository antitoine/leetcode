package main

import (
	"fmt"
	"math/bits"
)

func minimizeXor(num1 int, num2 int) int {
	bitCountNum1 := bits.OnesCount(uint(num1))
	bitCountNum2 := bits.OnesCount(uint(num2))
	for bitCountNum1 > bitCountNum2 {
		num1 &= num1 - 1
		bitCountNum1--
	}
	for bitCountNum1 < bitCountNum2 {
		lowestUnsetBit := 1
		for num1&lowestUnsetBit != 0 {
			lowestUnsetBit <<= 1
		}
		num1 |= lowestUnsetBit
		bitCountNum1++
	}
	return num1
}

func main() {
	num1 := 3
	num2 := 5
	result := minimizeXor(num1, num2)
	fmt.Printf("Result: %#v\n", result)
	if result == 3 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
