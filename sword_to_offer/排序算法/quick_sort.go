package sort

/*
** 快速排序
** 快速排序的思想是自上而下，它设定一个标杆，让其余元素依此标杆进行左右站队，然后在左右再次实行这个策略直到排序完成
** 快排不是稳定排序，时间复杂度通常为O(n*log n)，极低概率下退化为O(n2)，空间复杂度为O(0)，即可以原地置换
 */
func QuickSort(data []int) (list []int) {
	if len(data) == 1 {
		return
	}
	//设定第一个元素为标的
	obj := data[0]
	left, right := 1, len(data)-1
	for left < right {
		if data[right] >= obj { //在右边找到一个小于标的的值
			right--
		} else {
			if data[left] < obj { //在左边找到一个大于标的的值
				left++
			}
		}
		if data[left] > obj && data[right] <= obj { //两者属于标的两侧，则交换位置
			data[left], data[right] = data[right], data[left]
			right--
		}
	}
	if data[0]>data[left]{
		data[0], data[left] = data[left], data[0]
	}
	QuickSort(data[:left])
	QuickSort(data[right:])
	return data
}
