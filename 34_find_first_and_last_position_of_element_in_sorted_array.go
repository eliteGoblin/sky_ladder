package amazon

func searchRange(nums []int, target int) []int {
	start := lowerBoundIndex(nums, target)
	if start == len(nums) || nums[start] != target {
		return []int{-1, -1}
	}
	end := upperBoundIndex(nums, target)
	return []int{start, end - 1}
}
