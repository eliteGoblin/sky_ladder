class NumArray:

    def __init__(self, nums: List[int]):
        self.nums = nums
        self.presum = [0] * len(self.nums)

        for i in range(0, len(nums)):
            if i == 0:
                self.presum[0] = nums[0]
            else:
                self.presum[i] = self.presum[i-1] + nums[i]

    def get(self, i) -> int:
        if i < 0:
            return 0
        return self.nums[i]


    def sumRange(self, left: int, right: int) -> int:
        return self.presum[right] - self.presum[left-1]
        


# Your NumArray object will be instantiated and called as such:
# obj = NumArray(nums)
# param_1 = obj.sumRange(left,right)