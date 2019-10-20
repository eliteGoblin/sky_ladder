package amazon

func topKFrequent(nums []int, topK int) []int {
	mp := make(map[int]int)
	for _, v := range nums {
		mp[v]++
	}
	bucket := make([][]int, len(nums))
	for k, v := range mp {
		if bucket[v-1] == nil {
			bucket[v-1] = make([]int, 0, len(nums))
		}
		bucket[v-1] = append(bucket[v-1], k)
	}
	res := make([]int, 0, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		res = append(res, bucket[i]...)
		if len(res) >= topK {
			break
		}
	}
	return res[:topK]
}
