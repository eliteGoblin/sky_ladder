
def getModifiedArray(length: int, updates: List[List[int]]) -> List[int]:
    if length < 1:
        return []
    
    diff = [0] * length
    for update in updates:
        start, end, delta = update
        diff[start] += delta
        if end + 1 < length:
            diff[end + 1] -= delta

    res = [0] * length
    res[0] = diff[0]

    for i in range(1, length):
        res[i] += res[i-1] + diff[i]
    
    return res
