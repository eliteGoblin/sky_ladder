from __future__ import annotations

class Solution:
    def containsDuplicate(self, nums: List[int]) -> bool:
        st = set()
        for num in nums:
            if num in st:
                return True
            st.add(num)
        return False 

if __name__ == "__main__":
