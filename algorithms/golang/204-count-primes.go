package main

// https://leetcode-cn.com/problems/count-primes/
func countPrimes(n int) int {
	ht := make([]int, n+1)

	for i := 2; i*i < n; i++ {
		if ht[i] == 0 {
			for j := i * i; j <= n; j += i {
				ht[j] = 1
			}
		}
	}

	var count int
	for k := 2; k < n; k++ {
		if ht[k] == 0 {
			count++
		}
	}

	return count
}
