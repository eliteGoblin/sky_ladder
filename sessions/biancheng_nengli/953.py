
def is_word1_le_word2(word1: str, word2: str, dt: dict[int, str]) -> bool:
    if word1 == word2:
        return True
    i, j = 0, 0
    while i < len(word1) and j < len(word2):
        if word1[i] != word2[j]:
            return dt[word1[i]] < dt[word2[j]]
        i += 1
        j += 1
    if i == len(word1):
        return True
    return False

class Solution:
    def isAlienSorted(self, words: List[str], order: str) -> bool:
        dt = {}
        for i, e in enumerate(order):
            dt[e] = i
        
        pre_word = words[0]
        for i in range(1, len(words)):
            if not is_word1_le_word2(pre_word, words[i], dt):
                return False
            pre_word = words[i]
        
        return True
