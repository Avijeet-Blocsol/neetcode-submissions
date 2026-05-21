func merge(nums1 []int, m int, nums2 []int, n int) {
	l := len(nums1)

	for i := m - 1; i >= 0; i-- {
		nums1[i+(l-m)] = nums1[i]
	}

	a := l - m
	b := 0
	i := 0

	for a < l && b < n {
		if nums1[a] <= nums2[b] {
			nums1[i] = nums1[a]
			i += 1
			a += 1
		} else {
			nums1[i] = nums2[b]
			i += 1
			b += 1
		}
	}

	if b < n {
		for i := m + b; i < m+n; i++ {
			nums1[i] = nums2[b]
			b += 1
		}
	}
}
