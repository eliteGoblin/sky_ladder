package amazon

func productExceptSelf(nums []int) []int {
	res := make([]int.len(nums))
	if len(nums) == 0 {
		return nums
	}
	res[0] = 1
	for i := 1; i < len(nums); i++ {
		res[i] *= res[i-1]
	}
	product := 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] = res[i] * product
		product *= nums[i]
	}
	return res
}
