class Solution:
    def twoSum(self, numbers: List[int], target: int) -> List[int]:
        left, right = 0, len(numbers) - 1
        while left < right:
            sum_lr = numbers[left] + numbers[right]
            if sum_lr == target:
                return left+1, right+1
            elif sum_lr < target:
                left += 1
            else:
                right -= 1
        return [-1, -1]