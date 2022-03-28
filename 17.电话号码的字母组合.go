/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */

// @lc code=start
func letterCombinations(digits string) []string {
	var ans []string
	if len(digits) == 0 {
		return ans
	}
	tmp := [][]byte{{'a', 'b', 'c'}, {'d', 'e', 'f'}, {'g', 'h', 'i'}, {'j', 'k', 'l'}, {'m', 'n', 'o'}, {'p', 'q', 'r', 's'}, {'t', 'u', 'v'}, {'w', 'x', 'y', 'z'}}
	var path string
	var backTracking func(idx int)
	backTracking = func(idx int) {
		if len(path) == len(digits) {
			ans = append(ans, path)
			return
		}

		v := byte2int(digits[idx])
		for j := range tmp[v-2] {
			path += string(tmp[v-2][j])
			backTracking(idx + 1)
			path = path[:len(path)-1]
		}

	}
	backTracking(0)
	return ans
}

func byte2int(b byte) int {
	switch b {
	case '2':
		return 2
	case '3':
		return 3
	case '4':
		return 4
	case '5':
		return 5
	case '6':
		return 6
	case '7':
		return 7
	case '8':
		return 8
	case '9':
		return 9
	}
	return -1
}

// @lc code=end

