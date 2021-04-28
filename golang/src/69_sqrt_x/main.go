package _69_sqrt_x

func mySqrt(x int) int {
	if x <= 1 {
		return x
	}
	left, right := 0, x
	for left < right {
		mid := left + (right-left)/2
		if x/mid == mid {
			return mid
		}
		if x/mid < mid {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right - 1
}
