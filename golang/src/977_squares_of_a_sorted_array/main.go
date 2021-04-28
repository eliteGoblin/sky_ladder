package _977_squares_of_a_sorted_array

func sortedSquares(A []int) []int {
	var res []int
	i, j := 0, len(A)-1
	for {
		for i < len(A) && A[i] < 0 {
			i++
		}
		for j >= 0 && A[j] >= 0 {
			j--
		}
		if i >= len(A) || j < 0 {
			break
		}
		if A[i]*A[i] <= A[j]*A[j] {
			res = append(res, A[i]*A[i])
			i++
		} else {
			res = append(res, A[j]*A[j])
			j--
		}
	}
	for i < len(A) && A[i] >= 0 {
		res = append(res, A[i]*A[i])
		i++
	}
	for j >= 0 && A[j] < 0 {
		res = append(res, A[j]*A[j])
		j--
	}
	return res
}
