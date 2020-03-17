package p455

import (
	"sort"
)

// URL: https://leetcode-cn.com/problems/assign-cookies/

// FindContentChildren assign cookies
func FindContentChildren(g []int, s []int) int {
	// 将胃口值和尺寸值数组排序，这样只需要按顺序去分配饼干即可。
	sort.Ints(g)
	sort.Ints(s)

	j := 0
	length := len(g)
	for k := range s {
		if length <= j {
			break
		}
		// 由于已经排好序了，所以只要能满足胃口就行。
		// 因为我们能保证每一次分配的时候比当前需要分配饼干的孩子的胃口已经满足了，
		// 所以只需要考虑当前分配的饼干尺寸能满足孩子即可。
		if s[k] >= g[j] {
			j++
		}
	}

	return j
}
