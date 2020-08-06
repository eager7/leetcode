package findRepeatNumber

//使用二分法逐步缩减数组重复数字范围
func my(nums []int) int {
	start, end := 1, len(nums)
	for start != end {
		middle := (end-start)/2 + start       //找出中位数 23345
		if start == middle || end == middle { //这里需要特殊处理只剩有限相同数字的情况，如5,5,5,5,5,5,5
			return middle
		}

		var left, right int
		for _, n := range nums {
			if n > end || n < start {
				continue //跳过已经筛选过的数字
			}
			if n < middle {
				left++
			} else {
				right++
			}
		}
		if left >= right { //左边数量多，左边有重复
			end = middle
		} else {
			start = middle
		}
	}
	return start
}
