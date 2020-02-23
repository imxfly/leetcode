package main

// https://leetcode-cn.com/problems/first-unique-character-in-a-string/submissions/
func firstUniqChar(s string) int {
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
