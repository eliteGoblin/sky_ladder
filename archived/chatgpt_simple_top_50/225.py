from collections import deque

class MyStack:

    def __init__(self):
        self.q = deque()
        self.q
        self.tmp_q = deque()

    def push(self, x: int) -> None:
        self.q.appendleft(x)

    def pop(self) -> int:
        if len(self.q) == 0:
            return -1
        if len(self.q) == 1:
            return self.q.pop()
        
        while len(self.q) > 1:
            self.tmp_q.appendleft(self.q.pop())
        res = self.q.pop()
        self.q, self.tmp_q = self.tmp_q, self.q
        return res

    def top(self) -> int:
        res = self.pop()
        self.push(res)
        return res

    def empty(self) -> bool:
        return len(self.q) == 0
