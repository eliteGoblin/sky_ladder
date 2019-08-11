package amazon

import "sort"

func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}
	p := -1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			p = i
			break
		}
	}
	if p == -1 {
		sort.Ints(nums)
		return
	}
	s := p + 1
	for i := len(nums) - 1; i > p; i-- {
		if nums[i] > nums[p] {
			s = i
			break
		}
	}
	nums[p], nums[s] = nums[s], nums[p]
	sort.Ints(nums[p+1:])
}
