/*
 * @lc app=leetcode.cn id=674 lang=golang
 *
 * [674] 最长连续递增序列
 */

// @lc code=start
func findLengthOfLCIS(nums []int) int {
	dp := make([]int, 2)
	dp[0] = 1
	ans := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			dp[i%2] = dp[(i-1)%2] + 1
		} else {
			dp[i%2] = 1
		}
		if dp[i%2] > ans {
			ans = dp[i%2]
		}
	}
	return ans
}

// @lc code=end

