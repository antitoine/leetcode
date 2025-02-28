package main

import (
	"fmt"
)

func shortestCommonSupersequence(str1, str2 string) string {
	m, n := len(str1), len(str2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	result := make([]byte, 0, m+n)
	i, j := m, n
	for i > 0 || j > 0 {
		if i == 0 {
			j--
			result = append(result, str2[j])
		} else if j == 0 {
			i--
			result = append(result, str1[i])
		} else if dp[i][j] == dp[i-1][j] {
			i--
			result = append(result, str1[i])
		} else if dp[i][j] == dp[i][j-1] {
			j--
			result = append(result, str2[j])
		} else {
			i--
			j--
			result = append(result, str1[i])
		}
	}
	reverse(result)
	return string(result)
}

func reverse(arr []byte) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func main() {
	str1 := "abac"
	str2 := "cab"
	result := shortestCommonSupersequence(str1, str2)
	fmt.Printf("Result: %#v\n", result)
	if result == "cabac" {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
