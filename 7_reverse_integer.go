package amazon

const (
	uint32Max = ^uint32(0)
	int32Max  = int32(uint32Max >> 1)
)

func reverse(x int) int {
	var res int64
	for x > 0 {
		res *= 10
		res += int64(x) % 10
		x /= 10
	}
	if res > int64(int32Max) {
		return 0
	}
	return int(res)
}
