class Solution:
    def canMakeArithmeticProgression(self, arr: List[int]) -> bool:
        arr = sorted(arr)
        if len(arr) <= 2:
            return True
        interval: int = arr[1] - arr[0]
        for i in range(2, len(arr)):
            if arr[i] - arr[i-1] != interval:
                return False
        return True

