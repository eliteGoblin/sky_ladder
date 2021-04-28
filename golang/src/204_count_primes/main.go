package _204_count_primes

func countPrimes(n int) int {
	isPrime := make([]bool, n)
	for i := range isPrime {
		isPrime[i] = true
	}
	res := 0
	for i := 2; i < n; i++ {
		if !isPrime[i] {
			continue
		}
		res++
		for j := 2; i*j < n; j++ {
			isPrime[i*j] = false
		}
	}
	return res
}

// 素数筛:
//   初始认为所有数为素数
//   以2开始, 将2的倍数全标记为非素数
//   循环: 将后续素数的倍数全部即为非素数, 剩余没有标记的自然是素数
//   一遍遍的标记就像一次次的把非素数筛去
