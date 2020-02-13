package main

// https://leetcode-cn.com/problems/single-number/
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
