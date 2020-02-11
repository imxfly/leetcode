package main

/**
https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/

给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。

设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。

注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

示例 1:

输入: [7,1,5,3,6,4]
输出: 7
解释: 在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
     随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6-3 = 3 。
示例 2:

输入: [1,2,3,4,5]
输出: 4
解释: 在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
     注意你不能在第 1 天和第 2 天接连购买股票，之后再将它们卖出。
     因为这样属于同时参与了多笔交易，你必须在再次购买前出售掉之前的股票。
示例 3:

输入: [7,6,4,3,1]
输出: 0
解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。
*/
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
