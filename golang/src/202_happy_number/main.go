package _202_happy_number

func isHappy(n int) bool {
	st := make(map[int]struct{})
	for {
		st[n] = struct{}{}
		sum := 0
		for n > 0 {
			sum += (n % 10) * (n % 10)
			n /= 10
		}
		if sum == 1 {
			return true
		}
		if _, ok := st[sum]; ok {
			return false
		}
		n = sum
	}
}
