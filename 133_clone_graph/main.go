package _133_clone_graph

type Node struct {
      Val int
      Neighbors []*Node
}


func cloneGraph(node *Node) *Node {
	mp := make(map[*Node]*Node)
    return cloneNode(node, mp)
}

func cloneNode(node *Node, mp map[*Node]*Node)*Node{
	if node == nil {
		return nil
	}
	if _, ok := mp[node]; ok {
		return mp[node]
	}
	cNode := &Node{
		Val: node.Val,
	}
	mp[node] = cNode
	for _, e := range node.Neighbors {
		if e != nil {
			cNode.Neighbors = append(cNode.Neighbors, cloneNode(e, mp))
		}
	}
	return cNode
}