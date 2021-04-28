package _18_4sum

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	for i := 0; i < len(nums)-3; i++ {
		for i > 0 && nums[i] == nums[i-1] && i < len(nums)-3 {
			i++
		}
		for j := i + 1; j < len(nums)-2; j++ {
			for j > i+1 && nums[j] == nums[j-1] && j < len(nums)-2 {
				j++
			}
			k, l := j+1, len(nums)-1
			targetTwo := target - nums[i] - nums[j]
			for k < l {
				if nums[k]+nums[l] == targetTwo {
					res = append(res, []int{nums[i], nums[j], nums[k], nums[l]})
					for k < l {
						k++
						if nums[k] != nums[k-1] {
							break
						}
					}
				} else if nums[k]+nums[l] < targetTwo {
					k++
				} else {
					l--
				}
			}
		}
	}
	return res
}
