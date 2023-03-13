class MyQueue:

    def __init__(self):
        self.stk_tmp = []
        self.stk_final = []

    def push(self, x: int) -> None:
        while len(self.stk_final) > 0:
            self.stk_tmp.append(self.stk_final.pop())
        self.stk_tmp.append(x)
        while len(self.stk_tmp) > 0:
            self.stk_final.append(self.stk_tmp.pop())

    def pop(self) -> int:
        res = self.stk_final.pop()
        return res

    def peek(self) -> int:
        return self.stk_final[-1]

    def empty(self) -> bool:
        return len(self.stk_final) == 0


# Your MyQueue object will be instantiated and called as such:
# obj = MyQueue()
# obj.push(x)
# param_2 = obj.pop()
# param_3 = obj.peek()
# param_4 = obj.empty()