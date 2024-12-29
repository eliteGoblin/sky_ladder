class Solution:
    def reverse(self, x: int) -> int:
        is_pos = x >= 0
        x = x if is_pos else -x
        res = 0

        while x > 0:
            res = res * 10 + x % 10
            if res > 2**31 - 1:
                if not is_pos:
                    if res > 2**31:
                        return 0
                else:
                    return 0 
            x = x // 10

        return res if is_pos else -res
        

        

