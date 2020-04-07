package p349

// URL: https://leetcode-cn.com/problems/intersection-of-two-arrays/

// Intersection of two arrays
func Intersection(nums1 []int, nums2 []int) []int {
    ht := make(map[int]bool, len(nums1))
    for _, v := range nums1 {
        ht[v] = false
    }
    for _, v := range nums2 {
        if _, ok := ht[v]; ok {
            ht[v] = true
        }
    }

    var nums []int
    for k, v := range ht {
        if v {
            nums = append(nums, k)
        }
    }
    return nums
}
