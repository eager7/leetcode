package findMedianSortedArrays

/*
寻找中位数问题可以转换为将两个数组组合成一个数组的问题，并且保证组合的时间复杂度为log(m+n)
*/
func my(nums1 []int, nums2 []int) float64 {
	var temp1, tn2, total []int
	maxLen := len(nums1)
	if len(nums2) > maxLen {
		maxLen = len(nums2)
	}

	left, right := 0, 0
	for i := 0; i < maxLen; i++ {
		if nums1[i]<nums2[i] {
			total = append(total, nums1[i])
			temp1 = append(temp1, nums2[i])
		} else {
			left=nums1[i]
		}
	}
	return 0
}
