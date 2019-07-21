package amazon

import "sort"

func threeSum(nums []int) [][]int {
	var res [][]int
	if len(nums) < 3 {
		return res
	}
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		target := -nums[i]
		j, k := i+1, len(nums)-1
		for j < k {
			if j > i+1 && nums[j] == nums[j-1] {
				j++
				continue
			}
			v := nums[j] + nums[k]
			if v > target {
				k--
			} else if v < target {
				j++
			} else {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				j++
				k--
			}
		}
	}
	return res
}
