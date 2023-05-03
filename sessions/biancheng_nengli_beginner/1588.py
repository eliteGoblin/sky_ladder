class Solution:
    def sumOddLengthSubarrays(self, arr: List[int]) -> int:
        max_odd_sub_array_len = len(arr)
        if len(arr) % 2 == 0:
            max_odd_sub_array_len = len(arr) - 1
        
        res = 0
        for l in range(1, max_odd_sub_array_len+1, 2):
            for i in range(0, len(arr)):
                if i + l > len(arr):
                    break
                res += sum(arr[i:i+l])
        return res
