package amazon

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	midLeft := (len(nums1) + len(nums2) + 1) / 2
	midRight := (len(nums1) + len(nums2) + 2) / 2
	return float64(findKth(nums1, nums2, midLeft)+
		findKth(nums1, nums2, midRight)) / 2.0
}

func findKth(nums1, nums2 []int, k int) int {
	var s, l []int
	if len(nums1) <= len(nums2) {
		s, l = nums1, nums2
	} else {
		s, l = nums2, nums1
	}
	if len(s) == 0 {
		return l[k-1]
	}
	if k == 1 {
		if s[0] <= l[0] {
			return s[0]
		}
		return l[0]
	}
	if len(s) < k/2 {
		return findKth(s, l[k/2:], k-k/2)
	}
	if s[k/2-1] <= l[k/2-1] {
		return findKth(s[k/2:], l, k-k/2)
	}
	return findKth(s, l[k/2:], k-k/2)
}
