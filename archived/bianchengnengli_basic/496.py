class Solution:
    def nextGreaterElement(self, nums1: List[int], nums2: List[int]) -> List[int]:
        # find index mapping
        dt = {k:v for k, v in enumerate(nums1)}
        stk = []
        nums2_res = [0] * len(nums2)
        for i in range(len(nums2)-1, -1, -1):
            while len(stk) > 0 and stk[-1] < nums2[i]:
                stk.pop()
            nums2_res[i] = stk[-1] if len(stk) > 0 else -1
            stk.push(nums2[i])
        res = []
        for e in nums1:
            res.append(nums2_res[dt[e]])
        return res

