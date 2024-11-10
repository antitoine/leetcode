package main

import (
	"fmt"
	"math"
)

func minEnd(n int, x int) int64 {
	kMaxBit := int(math.Log2(float64(n)) + math.Log2(float64(x)) + 2)
	k := n - 1
	ans := int64(x)
	kBinaryIndex := 0

	for i := 0; i < kMaxBit; i++ {
		if (ans>>i)&1 == 0 {
			if (k>>kBinaryIndex)&1 == 1 {
				ans |= 1 << i
			}
			kBinaryIndex++
		}
	}

	return ans
}

func main() {
	n, x := 3, 4
	result := minEnd(n, x)
	fmt.Printf("Result: %#v\n", result)
	if result == 6 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
