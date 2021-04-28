package _16_3sum_closest

import "sort"

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	min := 999999
	for i := 0; i < len(nums)-2; i++ {
		target2 := target - nums[i]
		j, k := i+1, len(nums)-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if abs(target, sum) < abs(target, min) {
				min = sum
			}
			if nums[j]+nums[k] < target2 {
				j++
			} else if nums[j]+nums[k] == target2 {
				return target
			} else {
				k--
			}
		}
	}
	return min
}

func abs(a, b int) int {
	if a-b > 0 {
		return a - b
	}
	return b - a
}
