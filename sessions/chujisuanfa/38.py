

class Solution:
    def countAndSay(self, n: int) -> str:
        if n == 1:
            return "1"
        prev_say = self.countAndSay(n-1)
        # prev_say = "21"
        # 11 => 21
        count = 1
        res = ""
        for i in range(1, len(prev_say)):
            if prev_say[i] != prev_say[i-1]:
                res += str(count) + prev_say[i-1]
                count = 1
            else:
                count += 1
        if count > 0:
            res += str(count) + prev_say[-1]
        return res

s = Solution()

print(s.countAndSay(4))