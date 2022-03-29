/*
 * @lc app=leetcode.cn id=343 lang=golang
 *
 * [343] 整数拆分
 */

// @lc code=start
func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[2] = 1
	for i := 3; i <= n; i++ {
		for j := 1; j < i; j++ {
			dp[i] = max(dp[i], (i-j)*j)

			dp[i] = max(dp[i], dp[i-j]*j)
		}
	}
	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

