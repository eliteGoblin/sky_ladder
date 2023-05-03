# In function: nothing magical about input, just plain int
# Problem said "binary representation", and input 00000000000000000000000000001011, leetcode will convert it to plain int... so confusing
# input is a integer, binary representation is 1100110...(which base doesn't matter)
# return need to be integer, but underneath the binary representation need to be reverse of input


class Solution:
    def hammingWeight(self, n: int) -> int:
        res = 0
        while n > 0:
            if n % 2 == 1:
                res += 1
            n = n // 2
        return res