package p242

// URL: https://leetcode-cn.com/problems/valid-anagram/

// IsAnagram valid anagram
func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	ht := make(map[rune]int, len(s))
	// 考虑到 Unicode 字符集存在变长字节，所以需要分别对 s 和 t 进行循环遍历，不能合并
	for _, ch := range s {
		ht[ch]++
	}
	for _, ch := range t {
		ht[ch]--
	}
	for _, num := range ht {
		if num != 0 {
			return false
		}
	}

	return true
}
