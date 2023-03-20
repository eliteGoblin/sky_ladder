from enum import Enum

class State(Enum):
    START = 0
    CONSUME_DIGITS = 1

class Solution:
    def myAtoi(self, s: str) -> int:

        state = State.START
        is_positive = True
        res = 0

        for e in s:
            if state == state.START:
                if e == " ":
                    continue
                elif e.isdigit():
                    state = State.CONSUME_DIGITS
                    res = int(e)
                elif e == "-":
                    is_positive = False
                    state = State.CONSUME_DIGITS
                elif e == "+":
                    state = State.CONSUME_DIGITS
                else:
                    break
            elif state == state.CONSUME_DIGITS:
                if not e.isdigit():
                    break
                res = res * 10 + int(e)
                if is_positive:
                    if res >= (1 << 31) - 1:
                        return (1 << 31) - 1
                else:
                    if res >= (1 << 31):
                        return -1 * (1 << 31)
            else:
                assert False

        return res if is_positive else -res