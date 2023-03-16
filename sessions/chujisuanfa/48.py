class Solution:
    def rotate(self, matrix: List[List[int]]) -> None:
        """
        Do not return anything, modify matrix in-place instead.
        """
        # up switch with down row
        up, down = 0, len(matrix) - 1
        while up < down:
            matrix[up], matrix[down] = matrix[down], matrix[up]
            up += 1
            down -= 1
        
        # diag switch
        for i in range(0, len(matrix)):
            for j in range(0, i):
                matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
        