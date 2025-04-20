package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	weights := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&weights[i])
	}

	dp := make([]bool, m+1)
	dp[0] = true

	for _, weight := range weights {
		for j := m; j >= weight; j-- {
			if dp[j-weight] {
				dp[j] = true
			}
		}
	}

	for i := m; i >= 0; i-- {
		if dp[i] {
			fmt.Println(i)
			return
		}
	}
}
