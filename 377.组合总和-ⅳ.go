/*
 * @lc app=leetcode.cn id=377 lang=golang
 *
 * [377] 组合总和 Ⅳ
 */

// @lc code=start
func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for j := range nums {
			if i-nums[j] >= 0 {
				dp[i] += dp[i-nums[j]]
			}
		}
	}
	return dp[target]
}

// @lc code=end

