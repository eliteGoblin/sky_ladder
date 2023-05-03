import functools

def count_bits(e: int) -> int:
    res = 0
    while e != 0:
        if e % 2 == 1:
            res += 1
        e = e // 2
    return res

def my_compare(e1: int, e2: int) -> int:
    bits_e1 = count_bits(e1)
    bits_e2 = count_bits(e2)

    if bits_e1 != bits_e2:
        return bits_e1 - bits_e2
    
    return e1 - e2

class Solution:
    def sortByBits(self, arr: List[int]) -> List[int]:
        arr.sort(functools.cmp_to_key(my_compare))