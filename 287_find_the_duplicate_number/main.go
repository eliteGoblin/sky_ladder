package _287_find_the_duplicate_number

func findDuplicate(nums []int) int {
	slow, fast := 0, 0
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]

		if slow == fast {
			break
		}
	}
	start := 0
	for slow != start {
		slow = nums[slow]
		start = nums[start]
		if slow == start {
			break
		}
	}
	return start
}
