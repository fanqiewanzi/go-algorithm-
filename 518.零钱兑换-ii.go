/*
 * @lc app=leetcode.cn id=518 lang=golang
 *
 * [518] 零钱兑换 II
 */

// @lc code=start
func change(amount int, coins []int) int {
	// dp[i] 表示amount为i时有多少种方式凑成金额
	dp := make([]int, amount+1)
	dp[0] = 1
	for i := range coins {
		for j := coins[i]; j <= amount; j++ {
			dp[j] += dp[j-coins[i]]
		}
	}
	return dp[amount]
}

// @lc code=end

