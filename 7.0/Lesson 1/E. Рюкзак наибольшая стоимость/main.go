package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	weights := make([]int, n)
	costs := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&weights[i])
	}

	for i := 0; i < n; i++ {
		fmt.Scan(&costs[i])
	}

	dp := make([]int, m+1)

	for i := 0; i < n; i++ {
		for j := m; j >= weights[i]; j-- {
			dp[j] = max(dp[j], dp[j-weights[i]]+costs[i])
		}
	}

	fmt.Println(dp[m])
}
