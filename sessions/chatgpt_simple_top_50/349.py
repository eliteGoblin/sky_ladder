from typing import List

class Solution:
    def intersection(self, nums1: List[int], nums2: List[int]) -> List[int]:
        s1 = set(nums1)
        return list[set(e for e in nums2 if e in s1)]
    
s = Solution()
s.intersection([1,2,2,1], [2, 2])