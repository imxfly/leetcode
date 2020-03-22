package p496

import (
	"container/list"
	"fmt"
)

// NextGreaterElementFast next greater element
func NextGreaterElementFast(nums1 []int, nums2 []int) []int {
	result := make([]int, len(nums1))
	mMap := make(map[int]int)
	mStack := list.New()

	for _, v := range nums2 {
		if mStack.Len() == 0 {
			mStack.PushBack(v)
			continue
		}
		top, _ := mStack.Back().Value.(int)
		for v >= top {
			mMap[top] = v
			mStack.Remove(mStack.Back())
			if mStack.Len() == 0 {
				break
			}
			top, _ = mStack.Back().Value.(int)
		}
		mStack.PushBack(v)
	}

	fmt.Println(mMap)

	for i, v := range nums1 {
		item, ok := mMap[v]
		if !ok {
			result[i] = -1
		} else {
			result[i] = item
		}
	}

	return result
}

// NextGreaterElementSlow next greater element
func NextGreaterElementSlow(nums1 []int, nums2 []int) []int {
	result := make([]int, len(nums1))

	for nums1Key, nums1Value := range nums1 {
		result[nums1Key] = -1
		isStart := false
		for _, nums2Value := range nums2 {
			if nums2Value == nums1Value {
				isStart = true
				continue
			}
			if !isStart {
				continue
			}
			if isStart && nums2Value > nums1Value {
				result[nums1Key] = nums2Value
				break
			}
		}
	}

	return result
}
