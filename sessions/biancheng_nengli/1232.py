from typing import List

class Solution:
    def checkStraightLine(self, coordinates: List[List[int]]) -> bool:
        if len(coordinates) <= 1:
            return False
        if len(coordinates) == 2:
            return True
        
        p1 = coordinates[0]
        p2 = coordinates[1]

        if p1[0] == p2[0]:
            for i in range(2, len(coordinates)):
                if coordinates[i][0] != p1[0]:
                    return False
            return True
        
        ratio = (p2[1] - p1[1]) / (p2[0] - p1[0])

        for i in range(2, len(coordinates)):
            if coordinates[i][0] - p1[0] == 0:
                return False
            new_ratio = (coordinates[i][1] - p1[1]) / (coordinates[i][0] - p1[0])

            if abs(new_ratio - ratio) >= 1e-9:
                return False
        
        return True

s = Solution()

print(s.checkStraightLine([[1,2],[2,3],[3,4],[4,5],[5,6],[6,7],[1,8]]))