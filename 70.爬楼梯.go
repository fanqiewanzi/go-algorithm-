/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start
func climbStairs(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= 2; j++ {
			if i-j >= 0 {
				dp[i] += dp[i-j]
			}
		}
	}
	return dp[n]
}

// @lc code=end

