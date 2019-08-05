package amazon

func subarraySum(nums []int, k int) int {
	res := 0
	for i := range nums {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				res++
			}
		}
	}
	return res
}

func subarraySumMap(nums []int, k int) int {
	mp := make(map[int]int)
	mp[0] = 1
	sum := 0
	res := 0
	for i := range nums {
		sum += nums[i]
		res += mp[sum-k]
		mp[sum]++
	}
	return res
}
