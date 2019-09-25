package amazon

import "math/rand"

type Solution struct {
	sumArr []int
}

func Constructor(w []int) Solution {
	for i := 1; i < len(w); i++ {
		w[i] += w[i-1]
	}
	return Solution{
		sumArr: w,
	}
}

func (this *Solution) PickIndex() int {
	r := rand.Int() % this.sumArr[len(this.sumArr)-1]
	return upperBoundIndex(this.sumArr, r)
}
