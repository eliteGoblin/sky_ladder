package amazon

func findItinerary(tickets [][]string) []string {
	route := make(map[string]*myHeap)
	for i := range tickets {
		start := tickets[i][0]
		if _, ok := route[start]; !ok {
			route[start] = NewHeap(make([]interface{}, 0),
				func(x, y interface{}) bool {
					return x.(string) < y.(string)
				})
		}
		heap.Push(route[start], tickets[i][1]) // 切记不要写成route[start].Push, 调用heap.Push才会自动排序, 同理调用heap.Pop
	}
	res := make([]string, 0, len(tickets))
	dfsItinerary(route, "JFK", &res)
	i, j := 0, len(res) - 1
	for i < j {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}
	return res
}

func dfsItinerary(route map[string]*myHeap, cur string, res *[]string) {
	for route[cur] != nil && route[cur].Len() > 0 {
        nextElement := heap.Pop(route[cur])
        dfsItinerary(route, nextElement.(string), res)
	}
	*res = append(*res, cur)
}