
def lower_bound(nums: list[int], target: int) -> int:
    # 1 2 2 2 3
    left, right = 0, len(nums)

    while left < right:
        mid = left + (right - left) // 2
        value = nums[mid]

        if value == target:
            right = mid
        elif value < target:
            left = mid + 1
        elif value > target:
            right = mid
    
    return right

def upper_bound(nums: list[int], target: int) -> int:
    # 1 2 2 2 3
    left, right = 0, len(nums) # [left, right)

    while left < right:
        mid = left + (right - left) // 2
        value = nums[mid]

        if value == target:
            left = mid + 1
        elif value < target:
            left = mid + 1
        elif value > target:
            right = mid
    
    return left - 1

class Solution:
    def searchRange(self, nums: List[int], target: int) -> List[int]:
        if len(nums) == 0:
            return [-1, -1]
        if len(nums) == 1:
            if nums[0] == target:
                return [0, 0]
            return [-1, -1]

        if not (nums[0] <= target <= nums[-1]):
            return [-1, -1]
        
        return [lower_bound(nums, target), upper_bound(nums, target)]