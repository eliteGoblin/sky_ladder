class Difference:
    # 0 base
    def __init__(self, nums) -> None:
        self.diff = [0] * len(nums)
        self.diff[0] = nums[0]
    
    def incre(self, first, last, val) -> None:
        self.diff[first] += val
        if last + 1 < len(self.diff):
            self.diff[last+1] -= val
    
    def result(self) -> list[int]:
        res = [0] * len(self.diff)

        res[0] = self.diff[0]
        for i in range(1, len(self.diff)):
            res[i] = res[i-1] + self.diff[i]
        
        return res
    