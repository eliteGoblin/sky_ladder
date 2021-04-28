package _268_missing_number

func missingNumber(nums []int) int {
	allNums := make([]int, 0, len(nums)+1)
	allNums = append(allNums, nums...)
	allNums = append(allNums, -1)

	for i := 0; i < len(allNums); i++ {
		if allNums[i] == -1 {
			continue
		}
		for allNums[i] != i && allNums[i] != -1 {
			v := allNums[i]
			allNums[i], allNums[v] = allNums[v], allNums[i]
		}
	}

	for i := range allNums {
		if allNums[i] == -1 {
			return i
		}
	}
	return 0
}
