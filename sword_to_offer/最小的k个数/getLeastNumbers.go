package getLeastNumbers

func partition(arr []int, start, end int) int {
	if len(arr) == 0 {
		return 0
	}
	cur := start - 1 //游标
	for ; start < end; start++ {
		if arr[start] < arr[end] {
			cur++
			arr[cur], arr[start] = arr[start], arr[cur] //将小的数往前丢
		}
	}
	cur++
	arr[cur], arr[end] = arr[end], arr[cur]
	return cur //返回分界线
}

func getLeastNumbers(arr []int, k int) []int {
	var start, end = 0, len(arr)
	for {
		index := partition(arr, start, end)
		if index < k { //要找的点在右边
			start = index + 1
		} else if index > k {
			end = index - 1
		} else if index == k {
			return arr[:index]
		}
	}
}
