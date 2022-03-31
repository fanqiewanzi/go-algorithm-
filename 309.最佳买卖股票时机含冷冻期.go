/*
 * @lc app=leetcode.cn id=309 lang=golang
 *
 * [309] 最佳买卖股票时机含冷冻期
 */

// @lc code=start
func maxProfit(prices []int) int {
	// dp[i][0]表示第i天状态为买入股票的最大利润
	// dp[i][1]表示第i天状态为保持卖出股票的最大利润
	// dp[i][2]表示第i天状态为今天卖出股票的最大利润
	// dp[i][3]表示第i天状态为冷冻期的最大利润
	dp := make([][]int, 2)
	for i := range dp {
		dp[i] = make([]int, 4)
	}
	dp[0][0] = -prices[0]
	for i := 1; i < len(prices); i++ {
		// max(前一天买入，max(前一天为冷冻期，前一天保持卖出状态)-当前股票价格)
		dp[i%2][0] = max(dp[(i-1)%2][0], max(dp[(i-1)%2][3], dp[(i-1)%2][1])-prices[i])
		// max(前一天保持卖出状态，前一天为冷冻期)
		dp[i%2][1] = max(dp[(i-1)%2][1], dp[(i-1)%2][3])
		// 当前股票价格-买入价格,之前减过了所以这里直接加
		dp[i%2][2] = dp[(i-1)%2][0] + prices[i]
		// 前一天卖出股票的最大利润
		dp[i%2][3] = dp[(i-1)%2][2]
	}
	l := len(prices)
	return max(dp[(l-1)%2][2], max(dp[(l-1)%2][3], dp[(l-1)%2][1]))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

