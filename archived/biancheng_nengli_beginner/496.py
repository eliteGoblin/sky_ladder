class Solution:
    def nextGreaterElement(self, nums1: List[int], nums2: List[int]) -> List[int]:
        res = []
        for e in nums1:
            index = -1
            next_greater_index = -1

            for i in range(len(nums2)):
                if index == -1:
                    if nums2[i] == e:
                        index = i
                else:
                    if nums2[i] > nums2[index]:
                        next_greater_index = i
                        break
            if next_greater_index == -1:
                res.append(-1)
            else:
                res.append(nums2[next_greater_index])
        return res