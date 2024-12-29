# Write a function that takes the binary representation of an unsigned integer and returns the number of '1' bits it has (also known as the Hamming weight).
class Solution:
    def hammingWeight(self, n: int) -> int:
        res: int = 0
        while n != 0:
            # print(n)
            res += n & 1
            n >>= 1
        return res

s = Solution()
print(s.hammingWeight(-3))