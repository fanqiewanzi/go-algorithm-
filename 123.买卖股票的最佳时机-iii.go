/*
 * @lc app=leetcode.cn id=123 lang=golang
 *
 * [123] 买卖股票的最佳时机 III
 */

// @lc code=start
func maxProfit(prices []int) int {
	// dp[i][0]表示第i天什么都不做的最大利润
	// dp[i][1]表示第i天买入第一次的最大利润
	// dp[i][2]表示第i天卖出第一次的最大利润
	// dp[i][3]表示第i天买入第二次的最大利润
	// dp[i][4]表示第i天卖出第二次的最大利润
	dp := make([][]int, len(prices))
	for i := range dp {
		dp[i] = make([]int, 5)
	}
	// 初始化第0天的最大利润
	dp[0][1] = -prices[0]
	dp[0][3] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = dp[i-1][0]
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
		dp[i][2] = max(dp[i-1][1]+prices[i], dp[i-1][2])
		dp[i][3] = max(dp[i-1][3], dp[i-1][2]-prices[i])
		dp[i][4] = max(dp[i-1][4], dp[i-1][3]+prices[i])
	}
	return dp[len(dp)-1][4]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

