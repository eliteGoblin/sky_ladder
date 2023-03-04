class Solution:
    def subtractProductAndSum(self, n: int) -> int:
        digits: list[int] = []

        while n > 0:
            digits += n % 10
            n = n // 10
    
        product: int = 1
        for e in product:
            product *= e
        return product - sum(digits)
        