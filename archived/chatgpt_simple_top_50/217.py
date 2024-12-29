class Solution:
    def containsDuplicate(self, nums: List[int]) -> bool:
        st = set()
        for e in nums:
            if e in st:
                return False
            st.add(e)
        return True