from typing import List
import random

class Solution:
    def majorityElement(self, nums: List[int]) -> int:
        res = 0
        count = 0
        for e in nums:
            if count == 0:
                res = e
                count += 1
            elif e == res:
                count += 1
            else:
                count -= 1
        return res

# Kth smallest TLE
def partition(arr, low, high):
    pivot_index = random.randint(low, high)
    arr[pivot_index], arr[high] = arr[high], arr[pivot_index]
    pivot = arr[high]
    i = low - 1
    for j in range(low, high):
        if arr[j] <= pivot:
            i += 1
            arr[i], arr[j] = arr[j], arr[i]
    arr[i + 1], arr[high] = arr[high], arr[i + 1]
    return i + 1

def quick_select(arr, low, high, k):
    if low == high:
        return arr[low]
    
    pivot_index = partition(arr, low, high)
    if k == pivot_index:
        return arr[k]
    elif k < pivot_index:
        return quick_select(arr, low, pivot_index - 1, k)
    else:
        return quick_select(arr, pivot_index + 1, high, k)

def find_kth_smallest(arr, k):
    if k < 1 or k > len(arr):
        raise ValueError("k should be between 1 and the length of the input array")
    return quick_select(arr, 0, len(arr) - 1, k - 1)

class Solution:
    def majorityElement(self, nums: List[int]) -> int:
        nums = sorted(nums)
        return find_kth_smallest(nums, len(nums) // 2 + 1)
    

s = Solution()
print(s.majorityElement([3, 2, 3]))