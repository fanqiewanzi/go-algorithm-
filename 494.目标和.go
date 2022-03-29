/*
 * @lc app=leetcode.cn id=494 lang=golang
 *
 * [494] 目标和
 */

// @lc code=start
func findTargetSumWays(nums []int, target int) int {
	var sum int
	for i := range nums {
		sum += nums[i]
	}
	if (target+sum)%2 != 0 {
		return 0
	}
	s := (target + sum) / 2
	if s < 0 {
		s = -s
	}
	dp := make([]int, s+1)
	dp[0] = 1
	for i := range nums {
		for j := s; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
	}
	return dp[s]
}

// @lc code=end

