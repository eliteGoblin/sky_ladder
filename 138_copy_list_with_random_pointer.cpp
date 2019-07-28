class Solution {
public:
    Node* copyRandomList(Node* head) {
        unordered_map<Node*, Node*> m;
        return copyNode(head, m);
    }
    Node *copyNode(Node *node, unordered_map<Node *, Node *>& m) {
        if (!node) return nullptr;
        if (m.count(node)) return m[node];
        Node *res = new Node(node->val, nullptr, nullptr);
        m[node] = res;
        res -> next = copyNode(node->next, m);
        res -> random = copyNode(node->random, m);
        return res;
    }
};