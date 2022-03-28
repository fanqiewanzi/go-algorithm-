/*
 * @lc app=leetcode.cn id=968 lang=golang
 *
 * [968] 监控二叉树
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
func minCameraCover(root *TreeNode) int {
	var ans int
	// 1表示摄像头，2表示被覆盖，0表示需要增加摄像头
	var traverse func(node *TreeNode) int
	traverse = func(node *TreeNode) int {
		if node == nil {
			return 2
		}

		l := traverse(node.Left)
		r := traverse(node.Right)

		if l == 2 && r == 2 {
			return 0
		}

		if l == 0 || r == 0 {
			ans++
			return 1
		}

		if l == 1 || r == 1 {
			return 2
		}
		return -1
	}
	if traverse(root) == 0 {
		ans++
	}
	return ans
}

// @lc code=end

