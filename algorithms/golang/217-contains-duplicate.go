package main

// https://leetcode-cn.com/problems/contains-duplicate/
func containsDuplicate(nums []int) bool {
	// 用哈希表存储遍历过的数
    ht := make(map[int]bool)

    for _, num := range nums {
		// 如果在哈希表中存在，直接返回 true
        if _, ok := ht[num]; ok {
            return true
        }

		ht[num] = true
    }

	// 遍历完数组程序未结束说明每个数都不重复
    return false
}
