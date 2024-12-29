from header import *

class Solution:
    def rotate(self, matrix: List[List[int]]) -> None:
        """
        Do not return anything, modify matrix in-place instead.
        """
        # up down turn
        start_row, end_row = 0, len(matrix) - 1
        while start_row < end_row:
            matrix[start_row], matrix[end_row] = matrix[end_row], matrix[start_row]
            start_row += 1
            end_row -= 1
        # diag invert
        for i in range(len(matrix)):
            for j in range(0, i):
                matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
        