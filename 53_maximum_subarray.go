package amazon


func maxSubArray(nums []int) int {
	res := intMin
	curSum := intMin
	for _, v := range nums {
		curSum = max(curSum+v, v)
		res = max(res, curSum)
	}
	return res
}
