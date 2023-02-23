from header import *

class Solution:
    def deleteGreatestValue(self, grid: List[List[int]]) -> int:
        for row in grid:
            row.sort()
        res = 0
        for j in range(0, len(grid[0])):
            max_e = grid[0][j]
            for i in range(1, len(grid)):
                if grid[i][j] > max_e:
                    max_e = grid[i][j]
            res += max_e
        return res
