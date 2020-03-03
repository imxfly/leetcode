package main

func removeElement(nums []int, val int) int {
	i := 0
	last := len(nums)
	for i < last {
		if nums[i] == val {
			if nums[last-1] != val {
				nums[i] = nums[last-1]
			}
			last--
		} else {
			i++
		}
	}

	return last
}
