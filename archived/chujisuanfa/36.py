
def is_valid(lst: list[str]) -> bool:
    st = set()
    for e in lst:
        if e == ".":
            continue
        n = int(e)
        if n < 1 or n > 9:
            return False
        if n in st:
            return False
        st.add(n)
    return True

class Solution:
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        list_of_sequence_to_check: list[list[str]] = []

        for i in range(9):
            list_of_sequence_to_check.append(
                board[i]
            )
        
        for j in range(9):
            tmp: list[str] = []
            for i in range(9):
                tmp.append(board[i][j])
            list_of_sequence_to_check.append(tmp)
        
        for i in range(0, 9, 3):
            for j in range(0, 9, 3):
                tmp: list[str] = []
                for r in range(3):
                    for c in range(3):
                        tmp.append(board[i+r][j+c])
                list_of_sequence_to_check.append(tmp)
        
        for s in list_of_sequence_to_check:
            if not is_valid(s):
                return False
        
        return True
