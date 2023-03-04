from typing import List

class Solution:
    def average(self, salary: List[int]) -> float:
        if len(salary) == 0:
            return 0
        if len(salary) == 1:
            return salary[0]
        min_salary: int = salary[0]
        max_salary: int = salary[0]
        result: int = 0

        for i in range (0, len(salary)):
            min_salary = salary[i] if salary[i] < min_salary else min_salary
            max_salary = salary[i] if salary[i] > max_salary else max_salary
            result += salary[i]
        
        return (result - min_salary - max_salary) / (len(salary) - 2)

s = Solution()
print(s.average([4000,3000,1000,2000]))