from collections import defaultdict

class Solution:
    def canConstruct(self, ransomNote: str, magazine: str) -> bool:

        dtmag = defaultdict(int)

        for c in magazine:
            dtmag[c] += 1
        
        for c in ransomNote:
            if dtmag[c] == 0:
                return False
            dtmag[c] -= 1
        
        return True