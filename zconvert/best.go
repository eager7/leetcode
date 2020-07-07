package zconvert

/*
1. 不要把这个问题看成矩阵填充，否则会很麻烦
2. 想象顺序编译字符串时，字符会按照从上到下，在从下倒上的顺序移动，即0121 0121
3. 创建多个数组存储每行数据，然后遍历即可
*/
func best(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	var list = make([][]rune, numRows)

	var isr bool
	var cor int
	for _, c := range s {
		list[cor] = append(list[cor], c)
		if isr {
			cor--
		} else {
			cor++
		}
		if cor == numRows { //需折返
			cor = numRows - 2
			isr = true
		}
		if cor < 0 {
			cor = 1
			isr = false
		}
	}
	var ret string
	for i := 0; i < numRows; i++ {
		for _, c := range list[i] {
			ret += string(c)
		}
	}
	return ret
}
