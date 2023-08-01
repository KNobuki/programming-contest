package main

import "fmt"

func coinChange(coins []int, amount int) int {
	dp := make([][]int, len(coins)+1)
	for i := 0; i < len(coins)+1; i++ {
		dp[i] = make([]int, amount+1)
		for j := 0; j <= amount; j++ {
			dp[i][j] = -1
		}
	}
	dp[0][0] = 0
	for i := 0; i < len(coins); i++ {
		for j := 0; j <= amount; j++ {
			if dp[i][j] == -1 {
				continue
			}
			for k := 0; k*coins[i]+j <= amount; k++ {
				if dp[i+1][k*coins[i]+j] == -1 || dp[i][j]+k < dp[i+1][k*coins[i]+j] {
					dp[i+1][k*coins[i]+j] = dp[i][j] + k
				}
			}
		}
	}
	return dp[len(coins)][amount]
}

func main() {
	fmt.Println(coinChange([]int{1}, 0))
}
