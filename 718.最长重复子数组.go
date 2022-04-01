/*
 * @lc app=leetcode.cn id=718 lang=golang
 *
 * [718] 最长重复子数组
 */

// @lc code=start
// func findLength(nums1 []int, nums2 []int) int {
// 	dp := make([][]int, len(nums1)+1)
// 	for i := range dp {
// 		dp[i] = make([]int, len(nums2)+1)
// 	}
// 	var ans int
// 	for i := 1; i <= len(nums1); i++ {
// 		for j := 1; j <= len(nums2); j++ {
// 			if nums1[i-1] == nums2[j-1] {
// 				dp[i][j] = dp[i-1][j-1] + 1
// 			}
// 			if dp[i][j] > ans {
// 				ans = dp[i][j]
// 			}
// 		}
// 	}
// 	return ans
// }

func findLength(nums1 []int, nums2 []int) int {
	dp := make([]int, len(nums2)+1)
	var ans int
	for i := 1; i <= len(nums1); i++ {
		for j := len(nums2); j > 0; j-- {
			if nums1[i-1] == nums2[j-1] {
				dp[j] = dp[j-1] + 1
			} else {
				dp[j] = 0
			}
			if dp[j] > ans {
				ans = dp[j]
			}
		}
	}
	return ans
}

// @lc code=end

