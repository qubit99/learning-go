package main

import (
	"fmt"
	"strings"
)

func main() {
	var rows int
	fmt.Print("Enter the number of rows for Pascal's Triangle: ")
	fmt.Scan(&rows)

	var dp [][]int
	dp = append(dp, []int{1})
	dp = append(dp, []int{1, 1})

	for i := 2; i < rows; i++ {
		row := make([]int, i+1)
		row[0], row[i] = 1, 1
		for j := 1; j < i; j++ {
			row[j] = dp[i-1][j-1] + dp[i-1][j]
		}
		dp = append(dp, row)
	}

	for i := 0; i < rows; i++ {
		nums := make([]string, len(dp[i]))
		for j, val := range dp[i] {
			nums[j] = fmt.Sprintf("%d", val)
		}
		fmt.Println(strings.Repeat("  ", rows-i), strings.Join(nums, "   "))
	}
}
