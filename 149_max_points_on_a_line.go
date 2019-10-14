package A

func maxPoints(points [][]int) int {
	if len(points) <= 1 {
		return len(points)
	}
	res := 0
	for i := 0; i < len(points); i++ {
		dup := 0
		ct := make(map[[2]int]int)
		for j := i + 1; j < len(points); j++ {
			deltaX := points[i][0] - points[j][0]
			deltaY := points[i][1] - points[j][1]
			if deltaX == 0 && deltaY == 0 {
				dup++
				continue
			}
			gcd := getGCD(deltaX, deltaY)
			ct[[2]int{deltaX / gcd, deltaY / gcd}]++
		}
		if dup+1 > res {
			res = dup + 1
		}
		for _, v := range ct {
			if v+1+dup > res {
				res = v + dup + 1
			}
		}
	}
	return res
}

func getGCD(a, b int) int {
	if b == 0 {
		return a
	}
	return getGCD(b, a%b)
}
