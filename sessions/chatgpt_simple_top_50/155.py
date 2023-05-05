class MinStack:

    def __init__(self):
        self.mono_stk = []
        self.stk = []


    def push(self, val: int) -> None:
        self.stk.append(val)
        if len(self.mono_stk) == 0 or val <= self.mono_stk[-1]:
            self.mono_stk.append(val)

    def pop(self) -> None:
        val = self.stk.pop()
        if val == self.mono_stk[-1]:
            self.mono_stk.pop()

    def top(self) -> int:
        return self.stk[-1]

    def getMin(self) -> int:
        return self.mono_stk[-1]
