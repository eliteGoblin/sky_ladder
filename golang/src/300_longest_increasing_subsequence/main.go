package _300_longest_increasing_subsequence

// Memo方法: LISLenEndIn[i], 存放以i结尾的LIS, 必须包含nums[i]. 因此LISLenEndIn[last]并不是答案，还是遍历一遍，求max
//   注意findLISLenEndIn必须对全部index调用，不能只调用最后一个认为对全部调用.因为有逻辑:  `if nums[i] < nums[index] {`, 这样会漏掉一些情况求值.

func lengthOfLIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	LISLenEndIn := make(map[int]int)
	for i := range nums {
		findLISLenEndIn(nums, LISLenEndIn, i)
	}
	res := 0
	for _, v := range LISLenEndIn {
		if v > res {
			res = v
		}
	}
	return res
}

func findLISLenEndIn(nums []int, LISLenEndIn map[int]int, index int) int {
	if v, ok := LISLenEndIn[index]; ok {
		return v
	}
	res := 1
	for i := 0; i < index; i++ {
		if nums[i] < nums[index] {
			res = max(res, findLISLenEndIn(nums, LISLenEndIn, i)+1)
		}
	}
	LISLenEndIn[index] = res
	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
