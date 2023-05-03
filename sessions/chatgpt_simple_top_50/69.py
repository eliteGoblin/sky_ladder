class Solution:
    def mySqrt(self, x: int) -> int:
        if x <= 1:
            return x
        
        start, end = 0, x

        while start < end:
            mid = start + (end - start) // 2
            if mid * mid > x:
                end = mid
            else:
                start = mid + 1
        return end - 1