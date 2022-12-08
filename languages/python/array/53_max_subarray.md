
# containing at least one number


```py
class Solution:
def maxSubArray(self, nums: List[int]) -> int:
    max_sum_end_at, result = nums[0], nums[0]
    for e in nums[1:]:
        max_sum_end_at = max(max_sum_end_at + e, e)
        result = max(result, max_sum_end_at)
    return result
```

optimal sub structure:

max_sum_end_at[i] = max(max_sum_end_at[i-1] + nums[i], nums[i])

max_sum_end_at[index], 存放以index结尾的最大连续数组和; result为max_sum_end_at中最大的. 

因此初始化为: `max_sum_end_at, result = nums[0], nums[0]`

然后从index 1 开始遍历


# 其他解法

至少选择一种分情况讨论: 

*  全为负数: 至少选一个, 选最大的(负数绝对值最小的)
*  有>=0元素: 选正值(对结果有贡献的)


排除特殊情况: 全为负数: 变为可以不选元素: 

optimal sub structure 变为: 

max_sum_end_at[i] = max(max_sum_end_at[i-1] + nums[i], 0)

```py
class Solution:
  def maxSubArray(self, nums: List[int]) -> int:
      if max(nums) < 0:
          return max(nums)
      max_sum_end_at, result = max(nums[0], 0), max(nums[0], 0) # 防止nums[0] < 0, 不能选择
      for e in nums[1:]:
          max_sum_end_at = max(max_sum_end_at + e, 0)
          result = max(result, max_sum_end_at)
      return result
```

或者简化:

```py
class Solution:
  def maxSubArray(self, nums: List[int]) -> int:
      if max(nums) < 0:
          return max(nums)
      max_sum_end_at, result = 0, 0 # 初始状态谁也不选
      for e in nums:
          max_sum_end_at = max(max_sum_end_at + e, 0)
          result = max(result, max_sum_end_at)
      return result
```

# Kadane algorighm, 几个变种: 允许empty array, 不允许.
# https://en.wikipedia.org/wiki/Maximum_subarray_problem