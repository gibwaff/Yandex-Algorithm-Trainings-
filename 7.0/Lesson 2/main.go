package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var log2 []int
var stMax [][]int
var stCount [][]int

func buildSparseTable(arr []int) {
	n := len(arr)
	log2 = make([]int, n+1)
	for i := 2; i <= n; i++ {
		log2[i] = log2[i/2] + 1
	}

	K := log2[n] + 1
	stMax = make([][]int, n)
	stCount = make([][]int, n)
	for i := range arr {
		stMax[i] = make([]int, K)
		stCount[i] = make([]int, K)
		stMax[i][0] = arr[i]
		stCount[i][0] = 1
	}

	for j := 1; (1 << j) <= n; j++ {
		for i := 0; i+(1<<j) <= n; i++ {
			left := stMax[i][j-1]
			right := stMax[i+(1<<(j-1))][j-1]

			if left == right {
				stMax[i][j] = left
				stCount[i][j] = stCount[i][j-1] + stCount[i+(1<<(j-1))][j-1]
			} else if left > right {
				stMax[i][j] = left
				stCount[i][j] = stCount[i][j-1]
			} else {
				stMax[i][j] = right
				stCount[i][j] = stCount[i+(1<<(j-1))][j-1]
			}
		}
	}
}

func query(l, r int) (int, int) {
	length := r - l + 1
	j := log2[length]

	leftMax := stMax[l][j]
	rightMax := stMax[r-(1<<j)+1][j]

	if leftMax == rightMax {
		count := 0
		if stMax[l][j] == leftMax {
			count += stCount[l][j]
		}
		if stMax[r-(1<<j)+1][j] == rightMax {
			count += stCount[r-(1<<j)+1][j]
		}
		return leftMax, count
	} else if leftMax > rightMax {
		return leftMax, stCount[l][j]
	} else {
		return rightMax, stCount[r-(1<<j)+1][j]
	}
}

func main() {
	reader := bufio.NewScanner(os.Stdin)
	reader.Split(bufio.ScanWords)

	readInt := func() int {
		reader.Scan()
		val, _ := strconv.Atoi(reader.Text())
		return val
	}

	n := readInt()
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = readInt()
	}

	buildSparseTable(arr)

	k := readInt()
	for i := 0; i < k; i++ {
		l := readInt() - 1 // индекс с 0
		r := readInt() - 1
		maxVal, count := query(l, r)
		fmt.Println(maxVal, count)
	}
}
