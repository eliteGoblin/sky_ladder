class Solution:
    def dailyTemperatures(self, temperatures: List[int]) -> List[int]:
        stk_index = []
        temperatures_res = [-1] * len(temperatures)

        for i in range(len(temperatures) - 1, - 1, -1):
            while len(stk_index) > 0 and temperatures[i] >= temperatures[stk_index[-1]]:
                stk_index.pop()
            if len(stk_index) > 0:
                temperatures_res[i] = stk_index[-1]
            stk_index.append(i)
        
        res = []
        for i in range(len(stk_index)):
            tmp = temperatures_res[i] - i if temperatures_res[i] >= 0 else -1
        return res