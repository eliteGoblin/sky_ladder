package amazon

func search(nums []int, target int) int {
	return searchRecursive(nums, 0, len(nums)-1, target)
}

func searchRecursive(nums []int, start, last, target int) int {
	if start > last {
		return -1
	}
	mid := start + (last-start)/2
	if nums[mid] == target {
		return mid
	}
	if nums[mid] >= nums[start] {
		if target >= nums[start] && target < nums[mid] {
			return searchRecursive(nums, start, mid-1, target)
		}
		return searchRecursive(nums, mid+1, last, target)
	}
	if target > nums[mid] && target <= nums[last] {
		return searchRecursive(nums, mid+1, last, target)
	}
	return searchRecursive(nums, start, mid-1, target)
}
