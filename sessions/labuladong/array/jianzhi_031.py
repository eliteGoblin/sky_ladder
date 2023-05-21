class NumMatrix:

    def __init__(self, matrix: List[List[int]]):
        if len(matrix) == 0 or len(matrix[0]) == 0:
            return
        row = len(matrix)
        col = len(matrix[0])

        presum_at = self.presum_at
        self.presum = [[0] * col for _ in range(row)]

        print(self.presum)

        for i in range(row):
            for j in range(col):
                self.presum[i][j] = presum_at(i-1, j) + presum_at(i, j-1) + matrix[i][j] - presum_at(i-1, j-1)
        

    def presum_at(self, i, j):
        if i < 0 or j < 0:
            return 0
        return self.presum[i][j]

    def sumRegion(self, row1: int, col1: int, row2: int, col2: int) -> int:
        presum_at = self.presum_at
        return presum_at(row2, col2) - presum_at(row2, col1 -1) - presum_at(row1-1, col2) + presum_at(row1-1, col1-1)