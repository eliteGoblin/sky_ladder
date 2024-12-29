def square_sum(n) -> int:
    res = 0
    while n > 0:
        res = res + (n % 10) ** 2
        n = n // 2
    return res

class Solution:
    def isHappy(self, n: int) -> bool:
        if n == 1:
            return True
        
        tmp = set(n)

        while True:
            next = square_sum(n)
            if next == 1:
                return True
            if next in tmp:
                return False
            tmp.add(next)
            n = next

        