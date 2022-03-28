/*
 * @lc app=leetcode.cn id=763 lang=golang
 *
 * [763] 划分字母区间
 */

// @lc code=start
func partitionLabels(s string) []int {
	var mark [26]int
	for i := range s {
		mark[s[i]-'a'] = i
	}
	var ans []int
	var l, r int
	for i := range s {
		r = max(r, mark[s[i]-'a'])
		if i == r {
			ans = append(ans, r-l+1)
			l = i + 1
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

