package amazon

func findKthLargest(nums []int, k int) int {
	pos := partition(nums)
	if pos == k-1 {
		return nums[pos]
	}
	if pos > k-1 {
		return findKthLargest(nums[:pos], k)
	}
	return findKthLargest(nums[pos+1:], k-pos-1)
}

func partition(a []int) int {
	pivot := a[0]
	index := 0
	for i := 1; i < len(a); i++ {
		if a[i] >= pivot {
			index++
			a[index], a[i] = a[i], a[index]
		}
	}
	a[0], a[index] = a[index], a[0] // 之前没有这句，认为不需要换
	// 理由是: 左半边都是 >= pivot的，右半边都是 < pivot的，此规则成立，不需要换
	// case: [3,2,3,1,2,4,5,5,6], --> [3,3,4,5,5,6|1,2,2] index = 5, value = 6;
	// 左半边 >= 3; 右半边 < 3 但是partition的功能是: 使index左边都>=a[index]，这里a[index] = 6,破坏了约定，因此必须swap
	return index
}
