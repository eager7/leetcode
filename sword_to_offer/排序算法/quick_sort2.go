package sort

func Partition(data []int, start, end int) int {

	small := start - 1 //small表示比基准小的数字的游标
	for index := start; index < end; index++ {
		m, n := data[index], data[end]
		if m < n {
			small++             //游标向右移
			if small != index { //游标和index相等表示small和index同时递增，即没有遇到比基准大的值
				data[small], data[index] = data[index], data[small]
			}
		}
	}
	small++ //继续右移一位，把基准换进来
	data[small], data[end] = data[end], data[small]
	return small //返回分界线
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
