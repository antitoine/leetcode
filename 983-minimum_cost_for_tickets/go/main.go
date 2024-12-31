package main

import (
	"fmt"
)

func mincostTickets(days []int, costs []int) int {
	totalDays := len(days)
	lastDay := days[totalDays-1] + 1

	cost1DayPass, cost7DayPass, cost30DayPass := costs[0], costs[1], costs[2]

	dp := make([]int, lastDay)
	travelDays := make(map[int]bool)
	for _, day := range days {
		travelDays[day] = true
	}

	for day := 1; day < lastDay; day++ {
		if !travelDays[day] {
			dp[day] = dp[day-1]
			continue
		}
		costIfBuy1DayPass := dp[day-1] + cost1DayPass
		costIfBuy7DayPass := dp[max(0, day-7)] + cost7DayPass
		costIfBuy30DayPass := dp[max(0, day-30)] + cost30DayPass
		dp[day] = min(costIfBuy1DayPass, min(costIfBuy7DayPass, costIfBuy30DayPass))
	}

	return dp[lastDay-1]
}

func main() {
	days := []int{1, 4, 6, 7, 8, 20}
	costs := []int{2, 7, 15}
	result := mincostTickets(days, costs)
	fmt.Printf("Result: %#v\n", result)
	if result == 11 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
