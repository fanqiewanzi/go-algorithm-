/*
 * @lc app=leetcode.cn id=337 lang=golang
 *
 * [337] 打家劫舍 III
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rob(root *TreeNode) int {
	ans := robTree(root)
	return max(ans[0], ans[1])
}

func robTree(node *TreeNode) []int {
	if node == nil {
		return []int{0, 0}
	}
	left := robTree(node.Left)
	right := robTree(node.Right)

	val1 := node.Val + left[0] + right[0]
	val2 := max(left[0], left[1]) + max(right[0], right[1])
	return []int{val2, val1}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

