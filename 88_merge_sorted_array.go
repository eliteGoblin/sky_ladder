package amazon

// WA了N次，再次练习
func merge(nums1 []int, m int, nums2 []int, n int) {
	i1, i2, i := m-1, n-1, m+n-1
	for ; i1 >= 0 && i2 >= 0; i-- {
		if nums1[i1] >= nums2[i2] {
			nums1[i] = nums1[i1]
			i1--
		} else {
			nums1[i] = nums2[i2]
			i2--
		}
	}
    for i2 >=0 {
        nums1[i] = nums2[i2]
        i2--
        i--
    }
}