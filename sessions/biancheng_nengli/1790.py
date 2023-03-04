class Solution:
    def areAlmostEqual(self, s1: str, s2: str) -> bool:
        if s1 == s2:
            return True
        fst_diff_idx: int = -1
        sec_diff_idx: int = -1
        for idx, (x, y) in enumerate(zip(s1, s2)):
            if x != y:
                if fst_diff_idx < 0:
                    fst_diff_idx = idx
                elif sec_diff_idx < 0:
                    sec_diff_idx = idx
                else:
                    return False
        return s1[fst_diff_idx] == s2[sec_diff_idx] and s1[sec_diff_idx] == s2[fst_diff_idx]

s = Solution()

print(s.areAlmostEqual("bank", "kanb"))