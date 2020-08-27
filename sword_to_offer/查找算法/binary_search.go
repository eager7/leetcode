package search

//二分查找法，递归实现
func BinarySearch(list []int, target int) bool {
	if len(list) == 0 {
		return false
	}
	if len(list) == 1 {
		return list[0] == target
	}
	middle := list[len(list)>>1]
	if middle < target {
		return BinarySearch(list[len(list)>>1:], target)
	} else if middle > target {
		return BinarySearch(list[:len(list)>>1], target)
	} else {
		return true
	}
}

//二分查找法，非递归实现
func BinarySearch2(list []int, target int) bool {
	if len(list) == 0 {
		return false
	}
	le := len(list)
	start, middle, end := 0, le/2, le-1
	for start < middle {
		if list[middle] < target {
			start, middle = middle, middle+(end-middle)/2
		} else if middle > target {
			middle, end = middle-(middle-start)/2, middle
		} else {
			return true
		}
	}
	return list[start] == target
}
