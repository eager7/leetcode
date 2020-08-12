package min_array

func my(numbers []int) int {
	var min = numbers[0]
	for _, v := range numbers {
		if min > v {
			min = v
		}
	}
	return min
}

func best(numbers []int) int {
	if len(numbers) == 0 {
		panic("numbers is nil")
	}
	left, middle, right := 0, 0, len(numbers)-1
	for numbers[left] >= numbers[right] {
		if right-left == 1 {
			return numbers[right]
		}
		middle = (left + right) / 2
		if numbers[middle] == numbers[left] && numbers[middle] == numbers[right] { //遍历
			var min = numbers[left]
			for _, v := range numbers[left:right] {
				if min > v {
					min = v
				}
			}
			return min
		}
		if numbers[middle] >= numbers[left] { //前半部分递增
			left = middle //取后半部分
		} else if numbers[middle] <= numbers[right] { //后半部分递增
			right = middle //取前半部分
		}
	}
	return numbers[middle]
}
