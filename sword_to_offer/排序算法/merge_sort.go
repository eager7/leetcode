package sort

/*
** 归并排序
** 归并排序的思想是分而治之的思想，将数组按照二分法一直分割成1，2长度的小数组，然后各自排序后再组合
** 归并排序比较稳定，时间复杂度为O(n*log n)，空间复杂度为O(n)
 */
func MergeSort(data []int) (list []int) {
	if len(data) == 1 {
		return append(list, data[0])
	}
	if len(data) == 2 {
		if data[0] > data[1] {
			return append(list, data[1], data[0])
		}
		return append(list, data[0], data[1])
	}
	left, right := data[:len(data)/2], data[len(data)/2:]
	leftData, rightData := MergeSort(left), MergeSort(right)

	//归并结果，因为数组都是有序数组因此依次对比第一个值即可
	for len(leftData) > 0 && len(rightData) > 0 {
		if leftData[0] < rightData[0] {
			list = append(list, leftData[0])
			leftData = leftData[1:]
		} else {
			list = append(list, rightData[0])
			rightData = rightData[1:]
		}
	}
	return append(list, append(leftData, rightData...)...)//可能有剩余部分和一个空切片，都加进去即可
}
