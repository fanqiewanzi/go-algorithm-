/*
 * @lc app=leetcode.cn id=122 lang=golang
 *
 * [122] 买卖股票的最佳时机 II
 */

// @lc code=start
func maxProfit(prices []int) int {
	// dp[i][0]表示在第i天持有股票的最大现金
	// dp[i][1]表示在第i天不持有股票的最大现金
	dp := make([][]int, 2)
	for i := range dp {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i%2][0] = max(dp[(i-1)%2][0], dp[(i-1)%2][1]-prices[i])
		dp[i%2][1] = max(dp[(i-1)%2][1], prices[i]+dp[(i-1)%2][0])
	}

	return dp[(len(prices)-1)%2][1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

