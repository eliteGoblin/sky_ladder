from collections import defaultdict

class Solution:
    def intersect(self, nums1: List[int], nums2: List[int]) -> List[int]:
        dt = defaultdict(int)
        for e in nums1:
            dt[e] += 1
        
        res = []
        for e in nums2:
            if dt[e] > 0:
                res.append(e)
                dt[e] -= 1
        
        return res

