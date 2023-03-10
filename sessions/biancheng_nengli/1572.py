class Solution:
    def diagonalSum(self, mat: List[List[int]]) -> int:
        dup = 0
        if len(mat) % 2 == 1:
            mid = len(mat) // 2
            dup = mat[mid][mid]
        
        res = 0
        for i in range(len(mat)):
            res += mat[i][i]        # primary
            res += mat[i][len(mat) - i - 1]
        return res - dup