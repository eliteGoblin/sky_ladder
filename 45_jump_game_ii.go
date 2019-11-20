package amazon

func jump(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	pre, cur := -1, cur[0]
	steps := 1
	for cur < len(nums)-1 {
		next := cur
		for i := pre + 1; i <= cur; i++ {
			tmp := i + nums[i]
			if tmp > next {
				next = tmp
			}
		}
		pre = cur
		cur = next
		steps++
	}
	return steps
}
