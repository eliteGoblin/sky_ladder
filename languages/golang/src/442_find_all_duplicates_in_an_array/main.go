package _442_find_all_duplicates_in_an_array

func findDuplicates(nums []int) []int {
	res := make([]int, 0)
	for i := range nums {
		idx := abs(nums[i]) - 1
		if nums[idx] < 0 {
			res = append(res, -nums[i])
		}
		nums[idx] = -nums[idx]
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
