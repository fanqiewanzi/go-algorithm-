/*
 * @lc app=leetcode.cn id=63 lang=golang
 *
 * [63] 不同路径 II
 */

// @lc code=start
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := range dp {
		if obstacleGrid[i][0] == 1 {
			dp[i][0] = 0
			break
		} else {
			dp[i][0] = 1
		}
	}

	for j := range dp[0] {
		if obstacleGrid[0][j] == 1 {
			dp[0][j] = 0
			break
		} else {
			dp[0][j] = 1
		}
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}

// @lc code=end

