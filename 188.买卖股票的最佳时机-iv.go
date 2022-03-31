/*
 * @lc app=leetcode.cn id=188 lang=golang
 *
 * [188] 买卖股票的最佳时机 IV
 */

// @lc code=start
func maxProfit(k int, prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	dp := make([][]int, len(prices))
	for i := range dp {
		dp[i] = make([]int, k*2+1)
	}
	for i := 1; i < k*2+1; i += 2 {
		dp[0][i] = -prices[0]
	}

	for i := 1; i < len(prices); i++ {
		dp[i][0] = dp[i-1][0]
		for j := 1; j < 2*k+1; j++ {
			if j%2 == 0 {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-1]+prices[i])
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-1]-prices[i])
			}
		}
	}

	return dp[len(dp)-1][k*2]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

