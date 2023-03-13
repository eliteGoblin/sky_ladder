class Solution:
    def containsDuplicate(self, nums: List[int]) -> bool:
        st = set()
        for e in nums:
            if e in st:
                return True
            st.add(e)
        return False