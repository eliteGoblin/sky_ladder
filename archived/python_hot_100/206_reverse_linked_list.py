from typing import Optional, Any, Sequence, List
from header import *

class Solution:
    def reverseList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        fake_header = ListNode(next=head)
        p = fake_header.next
        while p is not None:
            next = p.next
            print(next)
            p.next = fake_header.next.next
            fake_header.next = p
            print(p.val, p.next)
            p = next
            break
        return fake_header.next


if __name__ == '__main__':
    n = ListNode(val=2)
    head = ListNode(val=1)
    s = Solution()
    new_head = s.reverseList(head)


    p = new_head
    while p is not None:
        print(f"({p.val}, {id(p.next)}")
        p = p.next