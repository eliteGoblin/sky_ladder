class Solution:
    def maximumWealth(self, accounts: List[List[int]]) -> int:
        res = 0
        for i in range(len(accounts)):
            row_sum = sum(accounts[i])
            res = row_sum if row_sum > res else res
        return res