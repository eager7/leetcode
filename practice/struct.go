package practice

type S1 struct {
	N int
	M int
}

type S2 struct {
	//M map[int]int // map类型不能比较大小
	//P []int // slice类型不能比较大小
	L  [2]int // 数组可以比较大小
	S1        // 可以比较的结构体可以比较大小
}
