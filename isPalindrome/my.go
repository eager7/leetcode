package isPalindrome

//拆分数字，判断是否是回文
func my(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	var list []int //拆分数字
	for x > 0 {
		y := x % 10
		list = append(list, y)
		x = x / 10
	}

	for i := 0; i < len(list)/2; i++ { //判断是否回文
		if list[i] != list[len(list)-i-1] {
			return false
		}
	}
	return true
}
