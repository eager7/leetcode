package findRepeatNumber

//使用抽屉原理将每个数字放到它下标的位置上，在放置过程中就会遇到重复数字
func my(nums []int) int {
	for i := range nums {
		for i != nums[i] {
			if nums[i] == nums[nums[i]] {
				return nums[i]
			}
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}
	return -1
}
