/*
 * @lc app=leetcode.cn id=406 lang=golang
 *
 * [406] 根据身高重建队列
 */

// @lc code=start
func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		a, b := people[i], people[j]
		return a[0] > b[0] || a[0] == b[0] && a[1] < b[1]
	})
	var ans [][]int
	for i := range people {
		idx := people[i][1]
		ans = append(ans[:idx], append([][]int{people[i]}, ans[idx:]...)...)
	}
	return ans
}

// @lc code=end

