package _349_intersection_of_two_arrays

func intersection(nums1 []int, nums2 []int) []int {
	mp1 := make(map[int]struct{})
	mp2 := make(map[int]struct{})
	for _, v := range nums1 {
		mp1[v] = struct{}{}
	}
	for _, v := range nums2 {
		mp2[v] = struct{}{}
	}
	res := make([]int, 0)
	for k := range mp1 {
		if _, ok := mp2[k]; ok {
			res = append(res, k)
		}
	}
	return res
}
