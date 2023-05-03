class Solution:
    def __init__(self):
        self.triangle = [[1], [1, 1], [1, 2, 1]]
        for i in range(3, 30):
            new_row = [1]
            for j in range(1, len(self.triangle[i-1])):
                new_row.append(self.triangle[i-1][j-1] + self.triangle[i-1][j])
            new_row.append(1)
            self.triangle.append(new_row)

    def generate(self, numRows: int) -> List[List[int]]:
        return self.triangle[:numRows]