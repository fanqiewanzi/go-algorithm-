/*
 * @lc app=leetcode.cn id=714 lang=golang
 *
 * [714] 买卖股票的最佳时机含手续费
 */

// @lc code=start
func maxProfit(prices []int, fee int) int {
	// dp[i][0] 表示拥有股票的最大利润
	// dp[i][1] 表示不拥有股票的最大利润
	dp := make([][]int, 2)
	for i := range dp {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i%2][0] = max(dp[(i-1)%2][0], dp[(i-1)%2][1]-prices[i])
		dp[i%2][1] = max(dp[(i-1)%2][1], dp[(i-1)%2][0]+prices[i]-fee)
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

