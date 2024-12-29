import sys

def man_dist(a: List[int], b: List[int]) -> int:
    return (a[0] - b[0] if a[0] > b[0] else b[0] - a[0] ) + \
            (a[1] - b[1] if a[1] > b[1] else b[1] - a[1])

class Solution:
    def nearestValidPoint(self, x: int, y: int, points: List[List[int]]) -> int:
        
        min_dist: int = sys.maxsize
        min_dist_ind: int = -1

        for i in range(0, len(points)):
            if x != points[i][0] and y != points[i][1]:
                continue
            dist: int = man_dist([x, y], points[i])
            if dist < min_dist:
                min_dist_ind = i
                min_dist = dist
        
        return min_dist_ind



