/*
 * @lc app=leetcode.cn id=714 lang=golang
 *
 * [714] 买卖股票的最佳时机含手续费
 */

// @lc code=start
func maxProfit(prices []int, fee int) int {
	var ans int
	minPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
		if prices[i] > minPrice+fee {
			ans += prices[i] - minPrice - fee
			minPrice = prices[i] - fee
		}
	}
	return ans
}

// @lc code=end

