package main

/**
https://leetcode-cn.com/problems/happy-number
编写一个算法来判断一个数是不是“快乐数”。

一个“快乐数”定义为：对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和，然后重复这个过程直到这个数变为 1，也可能是无限循环但始终变不到 1。如果可以变为 1，那么这个数就是快乐数。

示例:

输入: 19
输出: true
解释:
12 + 92 = 82
82 + 22 = 68
62 + 82 = 100
12 + 02 + 02 = 1
*/

// 用哈希表来解
func isHappy(n int) bool {
	var nums = make(map[int]bool)

	for n != 1 {
		if _, ok := nums[n]; ok {
			return false
		}
		nums[n] = true
		next := 0
		for n != 0 {
			next += (n % 10) * (n % 10)
			n /= 10
		}
		n = next
	}

	return true
}

// 注意到不管是不是快乐数，最终都会进入循环，所以可以使用「快慢指针」思想破循环
func calc(n int) int {
	sum := 0
	for n != 0 {
		sum += (n % 10) * (n % 10)
		n /= 10
	}
	return sum
}

func isHappy2(n int) bool {
	fast := n
	slow := n

	for {
		fast = calc(fast)
		slow = calc(slow)
		slow = calc(slow)
		if fast == slow {
			break
		}
	}

	return fast == 1
}
