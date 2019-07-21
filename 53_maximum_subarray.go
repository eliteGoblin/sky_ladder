package amazon

const (
	uintMax = ^uint(0)
	intMax  = int(uintMax >> 1)
	intMin  = -intMax - 1
)

func maxSubArray(nums []int) int {
	res := intMin
	curSum := intMin
	for _, v := range nums {
		curSum = max(curSum+v, v)
		res = max(res, curSum)
	}
	return res
}
