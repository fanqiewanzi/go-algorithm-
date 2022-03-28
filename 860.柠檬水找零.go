/*
 * @lc app=leetcode.cn id=860 lang=golang
 *
 * [860] 柠檬水找零
 */

// @lc code=start
func lemonadeChange(bills []int) bool {
	var five, ten int
	for i := range bills {
		if bills[i] == 5 {
			five++
		} else if bills[i] == 10 {
			if five > 0 {
				five--
				ten++
			} else {
				return false
			}
		} else if bills[i] == 20 {
			if ten > 0 && five > 0 {
				ten--
				five--
			} else if five >= 3 {
				five -= 3
			} else {
				return false
			}
		}
	}
	return true
}

// @lc code=end

