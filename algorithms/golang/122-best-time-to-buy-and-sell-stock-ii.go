package main

// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/
func maxProfit(prices []int) int {
	profit := 0
	length := len(prices) - 1
	for i := 0; i < length; i++ {
		// 因为我们得知了未来的股票价格，所以我们就不能错过每一次股票的上涨，
		// 即只要股票上涨就继续持有，股票下跌就马上卖出。
		// 这里的逻辑基础是：
		// 1. 连续增加的区间的变化值就是开始和结束的区间结点之差，所以频繁买入卖出的利润和在谷买入峰卖出的利润是一致的。
		// 2. 这里运用的就是贪心思想，即我要吃掉每一次涨跌幅的利润。
		if prices[i+1]-prices[i] > 0 {
			profit += prices[i+1] - prices[i]
		}
	}
	return profit
}
