package amazon

func twoSum(nums []int, target int) []int {
	mp := make(map[int]int)
	for i, v := range nums {
		mp[v] = i
	}
	for i, v := range nums {
		other := target - v
		if j, ok := mp[other]; ok {
			if i != j {
				return []int{i, j}
			}
		}
	}
	return []int{-1, -1}
}
