package p8

import (
	"strings"
)

// URL: https://leetcode-cn.com/problems/string-to-integer-atoi/

// MyAtoi convert string to int
func MyAtoi(str string) int {
	str = strings.Trim(str, " ")
	if len(str) < 1 {
		return 0
	}
	// 如果开头为正负号，保存并去掉正负号
	s0 := str[0]
	if s0 == '-' || s0 == '+' {
		str = str[1:]
		if len(str) < 1 {
			return 0
		}
	}

	n := 0
	// 循环每一个字节，直至出现非数字字符
	for i, ch := range []byte(str) {
		ch -= '0'
		if ch > 9 {
			break
		}
		n = n*10 + int(ch)

		// 如果数值超出范围，直接返回 INT_MAX 或者 INT_MIN
		if i > 8 && s0 == '-' && n > (1<<31) {
			return -(1 << 31)
		}
		if i > 8 && s0 != '-' && n >= (1<<31) {
			return (1 << 31) - 1
		}
	}

	if s0 == '-' {
		n = -n
	}
	return n
}
