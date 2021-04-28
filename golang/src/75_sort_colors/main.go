package _75_sort_colors

func sortColors(nums []int) {
	red, blue := -1, len(nums)
	for i := 0; i < blue; i++ {
		if nums[i] == 0 {
			red++
			nums[red], nums[i] = nums[i], nums[red]
		} else if nums[i] == 2 {
			blue--
			nums[blue], nums[i] = nums[i], nums[blue]
			i--
		}
	}
}
