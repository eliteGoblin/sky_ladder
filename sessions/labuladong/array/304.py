class NumMatrix:

    def __init__(self, matrix: List[List[int]]):
        if len(matrix) == 0:
            return
        m = len(matrix)
        n = len(matrix[0])
        self.presum = [[0] * n for _ in range(m)]

        at = self.psum

        for i in range(m):
            for j in range(n):
                self.presum[i][j] = at(i-1, j) + at(i, j-1) + matrix[i][j] - at(i-1, j-1)
    
    def psum(self, i, j) -> int:
        if i < 0 or j < 0:
            return 0
        return self.presum[i][j]

    def sumRegion(self, row1: int, col1: int, row2: int, col2: int) -> int:
        at = self.psum
        return at(row2, col2) - at(row2, col1-1) - at(row1-1, col2) + at(row1-1, col1-1)

