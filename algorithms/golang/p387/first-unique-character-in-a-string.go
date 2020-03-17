package p387

// URL: https://leetcode-cn.com/problems/first-unique-character-in-a-string/

// FirstUniqChar first unique character in a string
func FirstUniqChar(s string) int {
	ht := make(map[rune]int)

	for _, ch := range s {
		ht[ch]++
	}

	for i, ch := range s {
		if ht[ch] > 1 {
			continue
		}
		return i
	}

	return -1
}
