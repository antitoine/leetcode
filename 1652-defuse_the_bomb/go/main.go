package main

import (
	"fmt"
)

func decrypt(code []int, k int) []int {
	decryptedCode := make([]int, len(code))
	if k > 0 {
		for i := 0; i < len(code); i++ {
			sum := 0
			for j := i + 1; j <= i+k; j++ {
				sum += code[j%len(code)]
			}
			decryptedCode[i] = sum
		}
	} else if k < 0 {
		for i := len(code) - 1; i >= 0; i-- {
			sum := 0
			for j := i - 1; j >= i+k; j-- {
				if j >= 0 {
					sum += code[j]
				} else {
					sum += code[len(code)+j]
				}
			}
			decryptedCode[i] = sum
		}
	}
	return decryptedCode
}

func main() {
	code, k := []int{5, 7, 1, 4}, 3
	result := decrypt(code, k)
	fmt.Printf("Result: %#v\n", result)
	if fmt.Sprintf("%v", result) == "[12 10 16 13]" {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
