dt = {
    "I": 1,
    "V": 5,
    "X": 10,
    "L": 50,
    "C": 100,
    "D": 500,
    "M": 1000,
}

dt_2 = {
    "IV": 4,
    "IX": 9,
    "XL": 40,
    "XC": 90,
    "CD": 400,
    "CM": 900,
}

def next_token(s: str):
    next_2 = s[:2]
    if next_2 in dt_2:
        return dt_2[next_2], 2
    return dt[s[:1]], 1

class Solution:
    def romanToInt(self, s: str) -> int:
        index = 0
        res = 0

        while index < len(s):
            value, l = next_token(s[index:])
            res += value
            index += l
        
        return res
