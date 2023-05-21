
# class Solution:

#     def carPooling(self, trips: List[List[int]], capacity: int) -> bool:
#         # 1 base
#         diff = [0] * 1001
#         max_to = 0

#         for trip in trips:
#             count, f, to = trip
#             diff[f] += count
#             diff[to] -= count
            
#             max_to = to if to > max_to else max_to
        
#         cur_capacity = diff[0]
#         if cur_capacity > capacity:
#             return False
        
#         for i in range(1, max_to+1):
#             if diff[i] + cur_capacity > capacity:
#                 return False
#             cur_capacity += diff[i]
        
#         return True
        

from .array_template import Difference

class Solution:

    def carPooling(self, trips: List[List[int]], capacity: int) -> bool:

        diff = Difference([0] * 1001)

        for trip in trips:
            count, frm, to = trip

            diff.incre(frm, to-1, count)
        
        res = diff.result()

        for e in res:
            if e > capacity:
                return False
        return True