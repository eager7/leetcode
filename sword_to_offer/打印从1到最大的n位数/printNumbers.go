package printNumbers

import "fmt"

func printNumbers(n int) {
	var list = make([]int, 1)
	for len(list) <= n {
		for _, l := range list {
			fmt.Print(l)
		}
		fmt.Println()
		list = carry(list)
	}
}

func carry(list []int) []int {
	if len(list) == 0 {
		return append(make([]int, 0), 1) //这里新建slice，否则会覆盖旧数据
	}
	if list[len(list)-1] < 9 {
		list[len(list)-1] += 1 //不需进位直接加
		return list
	} else { //需要进位
		list[len(list)-1] = 0
		newList := carry(list[:len(list)-1]) //递归进高位
		return append(newList, list[len(list)-1])
	}
}
