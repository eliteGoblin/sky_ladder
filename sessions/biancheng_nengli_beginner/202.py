def sqaure_digits(n: int) -> int:
    res: int = 0
    while n > 0:
        res += (n % 10) ** 2
        n = n // 10
    return res

class Solution:
    def isHappy(self, n: int) -> bool:
        st = set()
        while True:
            n = sqaure_digits(n)
            if n == 1:
                return True
            if n in st:
                return False
            # print(n)
            st.add(n)
        return False

s = Solution()

print(s.isHappy(2))