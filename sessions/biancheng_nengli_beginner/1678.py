class Solution:
    def interpret(self, command: str) -> str:
        res = ""
        i = 0
        while i < len(command):
            if command[i:i+1] == "G":
                res += "G"
                i += 1
            elif command[i:i+2] == "()":
                res += "o"
                i += 2
            else:
                res += "al"
                i += 4
        
        return res
