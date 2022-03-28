/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	mp := make(map[int]int)
	for i := range nums {
		if v, ok := mp[target-nums[i]]; ok {
			return []int{i, v}
		} else {
			mp[nums[i]] = i
		}
	}
	return []int{}
}

// @lc code=end

