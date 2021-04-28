package _763_partition_labels

import "github.com/eliteGoblin/sky_ladder/common"

func partitionLabels(S string) []int {
	mp := make(map[byte]int)
	for i := range S {
		mp[S[i]] = i
	}
	res := make([]int, 0)
	for i := 0; i < len(S); {
		most := mp[S[i]]
		for j := i; j < most; j++ {
			most = common.Max(most, mp[S[j]])
		}
		res = append(res, most+1)
		i = most + 1
	}
	for i := len(res) - 1; i >= 1; i-- {
		res[i] = res[i] - res[i-1]
	}
	return res
}
