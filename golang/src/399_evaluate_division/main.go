package _399_evaluate_division

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	info := make(map[string]map[string]float64)
	for i := range equations {
		num, denum := equations[i][0], equations[i][1]
		if _, ok := info[num]; !ok {
			info[num] = make(map[string]float64)
		}
		info[num][denum] = values[i]
		if _, ok := info[denum]; !ok {
			info[denum] = make(map[string]float64)
		}
		info[denum][num] = 1.0 / values[i]
	}
	res := make([]float64, 0, len(queries))
	for i := range queries {
		visited := make(map[string]struct{})
		value := search(info, visited, queries[i][0], queries[i][1])
		res = append(res, value)
	}
	return res
}

func search(input map[string]map[string]float64, visited map[string]struct{}, num, denum string) float64 {
	if _, ok := input[num]; !ok {
		return -1.0
	}
	if _, ok := input[num][denum]; ok {
		return input[num][denum]
	}
	visited[num] = struct{}{}
	for k, v := range input[num] {
		if _, before := visited[k]; before {
			continue
		}
		visited[k] = struct{}{}
		d := search(input, visited, k, denum)
		if d > 0.0 {
			return d * v
		}
	}
	return -1.0
}
