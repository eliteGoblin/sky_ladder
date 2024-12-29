import queue 

class Solution:
    def matrixReshape(self, mat: List[List[int]], r: int, c: int) -> List[List[int]]:
        if len(mat) * len(mat[0]) != r * c:
            return mat
        my_q = queue.Queue()

        for i in range(len(mat)):
            for e in mat[i]:
                my_q.put(e)
        
        res = []

        for i in range(r):
            new_row = []
            for j in range(c):
                new_row.append(my_q.get())
            res.append(new_row)
        
        return res
        