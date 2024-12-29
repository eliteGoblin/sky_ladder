
dirs = [
    [0, 1], 
    [1, 0], 
    [0, -1],
    [-1, 0]
]

def next(mat, i, j, cur_dir):
    next_i, next_j = i + dirs[cur_dir][0], j + dirs[cur_dir][1]
    if next_i < 0 or next_i >= len(mat):
        return -1, -1
    if next_j < 0 or next_j >= len(mat[0]):
        return -1, -1
    return next_i, next_j

class Solution:
    def spiralOrder(self, matrix: List[List[int]]) -> List[int]:
        cur_dir = 0
        i, j = 0, 0
        res = []

        for ct in range(len(matrix) * len(matrix[0])):
            res.append(matrix[i][j])
            matrix[i][j] += 2000

            next_i, next_j = next(matrix, i, j, cur_dir)
            if next_i == -1 or abs(matrix[next_i][next_j]) > 1000:
                cur_dir = (cur_dir + 1) % 4
                next_i, next_j = next(matrix, i, j, cur_dir)
            i, j = next_i, next_j

        return res