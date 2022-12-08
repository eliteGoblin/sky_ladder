from __future__ import annotations
from collections import defaultdict

class Solution:
    def majorityElement(self, nums: List[int]) -> int:
        return moorVote(nums)

def moorVote(nums: List[int]) -> int:
	ct, res = 0, 0
	for num in nums:
		if ct == 0:
			ct += 1
			res = num
		else:
			ct = (ct + 1) if res == num else ct - 1
	return res
	
def naiveDict(nums: List[int]) -> int:
	d = defaultdict(int)
	
	for e in nums:
		d[e] = d[e] + 1
	return max(d, key=d.get)





if __name__ == "__main__":
	# print(majorityE([1, 2, 3, 3]))
	print(moorVote([1, 2, 3, 3, 3]))