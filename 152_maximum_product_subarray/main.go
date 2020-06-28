package _152_maximum_product_subarray

import "github.com/eliteGoblin/sky_ladder/common"

func maxProduct(nums []int) int {
	min := make([]int, len(nums))
	max := make([]int, len(nums))

	min[0] = nums[0]
	max[0] = nums[0]
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		min[i] = common.Min(
			common.Min(min[i-1]*nums[i], max[i-1]*nums[i]), nums[i])
		max[i] = common.Max(
			common.Max(min[i-1]*nums[i], max[i-1]*nums[i]), nums[i])
		if max[i] > res {
			res = max[i]
		}
	}
	return res
}
