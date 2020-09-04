package exchange

func exchange(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	cur, start, end := -1, 0, len(nums)

	for start < end {
		if nums[start]%2 != 0 { //奇数，往前丢
			cur++
			nums[cur], nums[start] = nums[start], nums[cur]
		}
		start++
	}
	return nums
}
