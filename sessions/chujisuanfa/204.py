
class Solution:
    def countPrimes(self, n: int) -> int:
        if n <= 1:
            return 0
        is_prime = [True] * n # index [1, n -1]
        """
        1 False
        2 True
        3 True
        """
        is_prime[1] = False
        res = 0 # 1 is 
        for i in range(2, n):
            if is_prime[i]:
                res += 1
            # mark i's x times til n
            for j in range(i+i, n, i):
                # use i*i instead of i+i, e.g: from 3, 3+3 been calculated in 2's round; 2 * 3; but 3*3 doesn't
                # any smaller multiple of i would have already been marked as composite by the previous primes.
                # i+i will TLE and i*i will pass
                is_prime[j] = False
        return res

