package main

import (
	"fmt"
)

func shortestSubarray(nums []int, k int) int {
	n := len(nums)
	prefixSum := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefixSum[i+1] = prefixSum[i] + nums[i]
	}

	var deque []int
	minLen := n + 1
	for i := 0; i <= n; i++ {
		for len(deque) > 0 && prefixSum[i]-prefixSum[deque[0]] >= k {
			minLen = min(minLen, i-deque[0])
			deque = deque[1:]
		}

		for len(deque) > 0 && prefixSum[i] <= prefixSum[deque[len(deque)-1]] {
			deque = deque[:len(deque)-1]
		}

		deque = append(deque, i)
	}

	if minLen == n+1 {
		return -1
	}
	return minLen
}

func main() {
	nums, k := []int{2, -1, 2}, 3
	result := shortestSubarray(nums, k)
	fmt.Printf("Result: %#v\n", result)
	if result == 3 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
