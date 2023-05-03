from icecream import ic
from typing import List

def rotate_matrix(mat):
    i, j = 0, len(mat)-1
    while i < j:
        mat[i], mat[j] = mat[j], mat[i]
        i += 1
        j -= 1
    for i in range(len(mat)):
        for j in range(i+1, len(mat)):
            mat[i][j], mat[j][i] = mat[j][i], mat[i][j]

def is_matrix_equal(mat, target):
    for i in range(len(mat)):
        for j in range(len(mat)):
            if mat[i][j] != target[i][j]:
                return False
    return True

class Solution:
    def findRotation(self, mat: List[List[int]], target: List[List[int]]) -> bool:
        if len(mat) == 1:
            return mat[0][0] == target[0][0]
        for i in range(4):
            if is_matrix_equal(mat, target):
                return True
            ic(mat)
            rotate_matrix(mat)
        return False

s = Solution()
res = s.findRotation(
    [[1,1],[0,1]],
    [[1,1],[1,0]]
)
ic(res)