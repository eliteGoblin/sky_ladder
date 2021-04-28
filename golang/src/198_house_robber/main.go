package _198_house_robber

import "github.com/eliteGoblin/sky_ladder/common"

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	memo := make([]int, len(nums))
	for i := range memo {
		memo[i] = -1
	}

	memo[0] = nums[0]
	memo[1] = common.Max(nums[0], nums[1])

	if len(nums) == 2 {
		return memo[1]
	}

	return recur(nums, len(nums)-1, memo)
}

func recur(nums []int, i int, memo []int) int {
	if memo[i] >= 0 {
		return memo[i]
	}
	memo[i] = common.Max(
		recur(nums, i-1, memo),
		recur(nums, i-2, memo)+nums[i])
	return memo[i]
}
