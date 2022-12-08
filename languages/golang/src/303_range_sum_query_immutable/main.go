package _303_range_sum_query_immutable

// [-2, 0, 3, -5, 2, -1]

type NumArray struct {
	dp []int
}

func Constructor(nums []int) NumArray {
	if len(nums) == 0 {
		return NumArray{}
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = dp[i-1] + nums[i]
	}
	return NumArray{
		dp: dp,
	}
}

func (this *NumArray) SumRange(i int, j int) int {
	if j >= len(this.dp) {
		return -1
	}
	if i == 0 {
		return this.dp[j]
	}
	return this.dp[j] - this.dp[i-1]
}
