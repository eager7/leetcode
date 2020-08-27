package sort

func Partition(data []int, start, end int) int {
	cur := start - 1 //small表示比基准小的数字的游标
	for index := start; index < end; index++ {
		if data[index] < data[end] {
			cur++             //游标向右移
			if cur != index { //游标和index相等表示small和index同时递增，即没有遇到比基准大的值
				data[cur], data[index] = data[index], data[cur]
			}
		}
	}
	cur++ //继续右移一位，把基准换进来
	data[cur], data[end] = data[end], data[cur]
	return cur //返回分界线
}

func QuickSort2(data[]int) {
	if len(data)==0{
		return
	}
	index := Partition(data, 0, len(data)-1)
	if len(data[:index]) > 0 {
		QuickSort2(data[:index])
	}
	if len(data[index+1:])>0{
		QuickSort2(data[index+1:])
	}
}
