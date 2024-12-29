
def is_arithmetic(nums) -> bool:
    if len(nums) < 2:
        return False
    nums = sorted(nums)
    diff = nums[1] - nums[0]
    for i in range(2, len(nums)):
        if nums[i] - nums[i-1] != diff:
            return False
    return True

class Solution:
    def checkArithmeticSubarrays(self, nums: List[int], l: List[int], r: List[int]) -> List[bool]:
        res = []
        for i in range(len(l)):
            sub_nums = nums[l[i]:r[i]+1]
            res.append(is_arithmetic(sub_nums))
        return res
