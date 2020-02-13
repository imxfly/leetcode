package main

// https://leetcode-cn.com/problems/happy-number

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
