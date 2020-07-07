package nreverse

func my(x int) int {
	var y = x
	if y < 0 {
		y = 0 - y //取绝对值
	}
	var store []int
	for y > 0 {
		store = append(store, y%10) //拆分位数
		y = y / 10                  //向左移位
	}

	//计算倍数
	multiple := func(n int) (m int) {
		m = 1
		for i := 0; i < n; i++ {
			m = m * 10
		}
		return m
	}
	//2147483647为32位最大值，重组数据不能大于这个数
	if len(store) == 10 {
		max := []int{2, 1, 4, 7, 4, 8, 3, 6, 4, 7}
		for i := 0; i < 10; i++ {
			if store[i] > max[i] {
				return 0
			} else if store[i] < max[i] {
				break
			}
		}
	}
	var out int
	for i := 0; i < len(store); i++ {
		out += store[i] * multiple(len(store)-i-1) //重组数据
	}
	if x < 0 {
		return 0 - out
	}
	return out
}
