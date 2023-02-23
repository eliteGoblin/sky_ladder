class Solution:
    def calculate_recur(self, candidates, cur_pos, cur_path, target, result):
        debug = False
        if len(cur_path) == 3 and cur_path[0] == 2 and cur_path[1] == 2:
            debug = True
        if debug:
            print(cur_pos, cur_path)
        if cur_pos >= len(candidates):
            return
        cur_path.append(candidates[cur_pos])
        path_sum = sum(cur_path)
        if path_sum > target:
            return
        elif path_sum == target:
            result.append(cur_path.copy())
            return

        for i in range(cur_pos, len(candidates)):
            if debug:
                print(i, len(candidates))
            self.calculate_recur(candidates, i, cur_path, target, result)
        cur_path.pop()

def combinationSum(self, candidates: List[int], target: int) -> List[List[int]]:
    result = []
    for i in range(0, len(candidates)):
        self.calculate_recur(candidates, i, [], target, result)
    return result