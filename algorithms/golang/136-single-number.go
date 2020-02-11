package main

import "fmt"

/*
https://leetcode-cn.com/problems/single-number/
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。

说明：

你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？

示例 1:

输入: [2,2,1]
输出: 1
示例 2:

输入: [4,1,2,1,2]
输出: 4
*/
func singleNumber(nums []int) int {
	var num2cnt = make(map[int]int)

	for _, num := range nums {
		if _, ok := num2cnt[num]; ok {
			delete(num2cnt, num)
		} else {
			num2cnt[num] = 1
		}
	}

	for num := range num2cnt {
		return num
	}

	return 0
}
