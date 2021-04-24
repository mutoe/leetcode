package best_time_to_buy_and_sell_stock_ii

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii

// level: 1
// time: O(n) 4ms 94.38%
// space: O(1) 3.4M 6.02%

//leetcode submit region begin(Prohibit modification and deletion)
func maxProfit(prices []int) int {
	var sum int
	for i := 0; i < len(prices)-1; i++ {
		if prices[i] < prices[i+1] {
			sum += prices[i+1] - prices[i]
		}
	}

	return sum
}

//leetcode submit region end(Prohibit modification and deletion)
