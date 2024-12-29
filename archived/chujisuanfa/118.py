cache = [[1]]

def generate_recur(row: int) -> None:
    if row < len(cache):
        return
    
    generate_recur(row - 1)

    new_row = [cache[row-1][0], cache[row-1][-1]]
    for i in range(1, len(cache[row-1])):
        new_row.insert(1, cache[row-1][i-1] + cache[row-1][i])
    
    cache.append(new_row)

    
class Solution:
    def generate(self, numRows: int) -> List[List[int]]:
        generate_recur(numRows - 1)
        return cache[:numRows]