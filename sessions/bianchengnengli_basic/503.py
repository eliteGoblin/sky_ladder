class Solution:
    def nextGreaterElements(self, nums: List[int]) -> List[int]:
        stk = []
        len_nums = len(nums)
        nums = nums * 2
        res = []

        for i in range(len(nums)-1, -1, -1):
            while len(stk) > 0 and stk[-1] <= nums[i]:
                stk.pop()
            res.insert(0, stk[-1] if len(stk) > 0 else -1)
            stk.append(nums[i])
        
        return res[:len_nums]
            