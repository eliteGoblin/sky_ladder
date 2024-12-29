class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        if len(prices) == 1:
            return 0
        
        max_profits = 0
        cur_min = prices[0]
        for i in range(1, len(prices)):
            if prices[i] - cur_min > max_profits:
                max_profits = prices[i] - cur_min
            if prices[i] < cur_min:
                cur_min = prices[i]
        
        return max_profits
