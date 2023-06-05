class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution:
    def mergeTwoLists(self, list1: Optional[ListNode], list2: Optional[ListNode]) -> Optional[ListNode]:
        pseudo_head = ListNode()
        pre = pseudo_head

        while list1 is not None and list2 is not None:
            if list1.val <= list2.val:
                pre = pre.next
                list1 = list1.next
            else:
                pre.next = list2
                list2 = list2.next
            pre = pre.next
        
        if list1 is not None:
            pre.next = list1
        if list2 is not None:
            pre.next = list2
        
        return pseudo_head.next