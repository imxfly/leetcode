package main

// https://leetcode-cn.com/problems/remove-element/
func removeElement(nums []int, val int) int {
	i := 0
	last := len(nums)

	for i < last {
		if nums[i] == val {
			nums[i] = nums[last-1]
			last--
		} else {
			i++
		}
	}

	return last
}
