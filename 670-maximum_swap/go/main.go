package main

import (
	"fmt"
	"strconv"
)

func maximumSwap(num int) int {
	nbDigits := len(strconv.Itoa(num))
	if nbDigits <= 1 {
		return num
	}
	biggestDigit := 0
	digits := make([]int, nbDigits)
	for i := nbDigits - 1; i >= 0; i-- {
		remainder := num % 10
		digits[i] = remainder
		num /= 10
		if biggestDigit < digits[i] {
			biggestDigit = digits[i]
		}
	}
	start := 0
	for start < nbDigits {
		if biggestDigit == digits[start] {
			start++
		} else {
			refDigit := digits[start]
			found := false
			for j := start + 1; j < nbDigits; j++ {
				if digits[j] > refDigit {
					found = true
					break
				}
			}
			if found {
				break
			}
			biggestDigit = digits[start]
			start++
		}
	}
	for i := start + 1; i < nbDigits; i++ {
		if digits[start] < digits[i] {
			biggestIdx := i
			for j := i + 1; j < nbDigits; j++ {
				if digits[biggestIdx] <= digits[j] {
					biggestIdx = j
				}
			}
			digits[start], digits[biggestIdx] = digits[biggestIdx], digits[start]
			break
		}
	}
	var result string
	for _, digit := range digits {
		result += strconv.Itoa(digit)
	}
	newNum, _ := strconv.Atoi(result)
	return newNum
}

func main() {
	num := 2736
	result := maximumSwap(num)
	fmt.Printf("Result: %#v\n", result)
	if result == 7236 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
