import heapq
from icecream import ic
from typing import List

# import math

# class Solution:
#     def kClosest(self, points: List[List[int]], k: int) -> List[List[int]]:
#         points = sorted(points, key=lambda x: math.sqrt(x[0]*x[0] + x[1]*x[1]))
#         return points[:k]
    

class SolutionHeap:
    def kClosest(self, points: List[List[int]], k: int) -> List[List[int]]:
        q = [(x**2 + y**2, i) for i, (x, y) in enumerate(points[:k])]
        heapq.heapify(q)
        for i in range(k, len(points)):
            x, y = points[i]
            heapq.heappushpop(q, (x**2, y**2, i))
            ic(q)
        return [points[index] for _, index in q]

s = SolutionHeap()
ic(s.kClosest([[1,3],[-2,2]], 1))