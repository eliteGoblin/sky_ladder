
class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        prefix_array = [1] * len(nums)
        for i in range(1, len(nums)):
            prefix_array[i] = prefix_array[i-1] * nums[i-1]
        suffix_array = [1] * len(nums)
        for j in range(len(nums)-2, -1, -1):
            suffix_array[j] = suffix_array[j+1] * nums[j+1]
        answers = []
        for i in range(0, len(nums)):
            answers.append(prefix_array[i] * suffix_array[i])
        return answers
