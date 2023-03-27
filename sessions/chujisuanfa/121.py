# Recursive will TLE at extreme case
# class Solution:
#     def __init__(self):
#         self.cache = {}
#         self.cache[0] = 0 # cache[i] means sell at ith day, max profit

#     def max_profit_recur(self, prices, j) -> int:
#         max_profit_end_day_j = 0

#         if j in self.cache:
#             return self.cache[j]
#         if j not in self.cache:
#             for i in range(j):
#                 profit = prices[j] - prices[i] + self.max_profit_recur(prices, i)
#                 if profit > max_profit_end_day_j:
#                     max_profit_end_day_j = profit

#         self.cache[j] = max_profit_end_day_j
#         return max_profit_end_day_j
        
#     def maxProfit(self, prices: List[int]) -> int:
#         self.max_profit_recur(prices, len(prices) - 1) # populate cache
#         res = 0
#         for i in range(len(prices)):
#             if self.cache[i] > res:
#                 res = self.cache[i]
#         return res

class Solution:

    def maxProfit(self, prices: List[int]) -> int:
        min_price = prices[0]
        max_profit = 0

        for i in range(1, len(prices)):
            min_price = min(prices[i], min_price)
            max_profit = max(max_profit, prices[i] - min_price)
        
        return max_profit

