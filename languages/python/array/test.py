from typing import List


class Solution:
    def sortedSquares(self, nums: List[int]) -> List[int]:
        fst_pos = len(nums)
        for idx, e in enumerate(nums):
            if e >= 0:
                fst_pos = idx
                break

        np, pp = fst_pos - 1, fst_pos
        res = []
        while np >= 0 and pp < len(nums):
            if nums[np] ** 2 <= nums[pp] ** 2:
                res.append(nums[np] ** 2)
                np -= 1
            else:
                res.append(nums[pp] ** 2)
                pp += 1
        while np >= 0:
            res.append(nums[np] ** 2)
            np -= 1
        while pp < len(nums):
            res.append(nums[pp] ** 2)
            pp += 1
        return res


if __name__ == "__main__":
    s = Solution()
    s.sortedSquares([-4,-1,0,3,10])
