from icecream import ic

class Solution:
    def nextGreaterElement(self, n: int) -> int:
        s_n = str(n)
        # find first asc position
        pos = -1

        for i in range(len(s_n)-1, 0, -1):
            if s_n[i] > s_n[i-1]:
                pos = i - 1
                break

        if pos == -1:
            return -1
        # find next greater position
        next_greater = pos + 1
        for i in range(next_greater, len(s_n)):
            if (s_n[i] > s_n[pos]) and s_n[i] < s_n[next_greater]:
                next_greater = i
        # swap pos and next_greater
        l = list(s_n)

        l[pos], l[next_greater] = l[next_greater], l[pos]

        ic(sorted(l[pos+1:]))
        new_l = l[:pos+1]
        new_l.extend(sorted(l[pos+1:]))
        ic(new_l)

        res = int(''.join(new_l))
        if res > 2**32:
            return -1

s = Solution()

ic(s.nextGreaterElement("230241"))