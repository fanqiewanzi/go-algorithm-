/*
 * @lc app=leetcode.cn id=416 lang=golang
 *
 * [416] 分割等和子集
 */

// @lc code=start
func canPartition(nums []int) bool {
	var sum int
	for i := range nums {
		sum += nums[i]
	}

	if sum%2 != 0 {
		return false
	}
	target := sum >> 1
	dp := make([]int, target+1)
	for i := 0; i < len(nums); i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] = max(dp[j], dp[j-nums[i]]+nums[i])
		}
	}
	if dp[target] == target {
		return true
	}
	return false
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

