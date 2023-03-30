from icecream import ic
from typing import List

import random

class Solution:
    nums: List[int]

    def __init__(self, nums: List[int]):
        self.nums = nums

    def reset(self) -> List[int]:
        return self.nums

    def shuffle(self) -> List[int]:
        tmp = self.nums[:]
        res = []
        while len(tmp) > 0:
            item = random.choice(tmp)
            tmp.remove(item)
            res.append(item)
        return res


def generate_permutation(nums: List[int]) -> List[List[int]]:
    if len(nums) == 1:
        return [nums]
    if len(nums) == 0:
        return []
    
    nums_less_1_res = generate_permutation(nums[:len(nums)-1])
    res = []
    for lst in nums_less_1_res:
        for i in range(len(lst)+1):
            clone = lst[:]
            clone.insert(i, nums[-1])
            res.append(clone)
    return res


class SolutionPermutation:
    # TLE, 13 items, this algorithm is too slow, seems not applicable for solution
    nums: List[int]
    cache: List[List[int]]

    def get_cache(self) -> List[List[int]]:
        if self.cache is None:
            self.cache = generate_permutation(self.nums)
        return self.cache

    def __init__(self, nums: List[int]):
        self.nums = nums
        self.get_cache()

    def reset(self) -> List[int]:
        return self.nums

    def shuffle(self) -> List[int]:
        cache = self.get_cache()
        return cache[
            random.randint(0, len(cache))
        ]

class SolutionSwap:

    def __init__(self, nums: List[int]):
        self.nums = nums

    def reset(self) -> List[int]:
        return self.nums


    def shuffle(self) -> List[int]:
        res = self.nums[:]

        for index in range(len(res)-1, -1, -1):
            rand_pos = random.randint(0, index)
            res[index], res[rand_pos] = res[rand_pos], res[index]
        
        return res



# Your Solution object will be instantiated and called as such:
# obj = Solution(nums)
# param_1 = obj.reset()
# param_2 = obj.shuffle()