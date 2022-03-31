/*
 * @lc app=leetcode.cn id=213 lang=golang
 *
 * [213] 打家劫舍 II
 */

// @lc code=start
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	l := robTracking(nums, 0, len(nums)-2)
	r := robTracking(nums, 1, len(nums)-1)
	return max(l, r)
}

func robTracking(nums []int, left, right int) int {
	if left == right {
		return nums[left]
	}
	dp := make([]int, len(nums))
	dp[left] = nums[left]
	dp[left+1] = max(nums[left], nums[left+1])
	for i := left + 2; i <= right; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[right]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

