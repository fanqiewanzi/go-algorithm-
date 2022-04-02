/*
 * @lc app=leetcode.cn id=115 lang=golang
 *
 * [115] 不同的子序列
 */

// @lc code=start
func numDistinct(s string, t string) int {
	// dp[i][j] 表示从0到i-1和0到j-1出现t的个数
	dp := make([][]int, len(s)+1)
	for i := range dp {
		dp[i] = make([]int, len(t)+1)
		dp[i][0] = 1
	}

	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(t); j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[len(s)][len(t)]

}

// @lc code=end

