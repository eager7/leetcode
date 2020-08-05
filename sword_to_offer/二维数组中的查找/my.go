package findNumberIn2DArray

/*
** 数组是有序的，从左上角到右下角递增，
** 因此可以通过遍历数组的右下角边对数组进行删减，从而减少遍历的量级
 */
func my(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	n, m := len(matrix), len(matrix[0])

	//当矩阵数组的第一个数大于目标值，后续部分一定都大于目标值，
	//同理，当最后一个数小于目标值，那么前导部分也一定小于目标值
	startN, endN := 0, n
	for i := 0; i < n; i++ {
		if matrix[i][m-1] < target { //裁剪掉多余部分
			startN = i+1 //当前列丢掉
		}
		if matrix[i][0] > target {
			endN = i //当前列丢掉
			break
		}
	}
	if startN >= endN {
		return false
	}
	//重新计算数组大小
	matrix = matrix[startN:endN]
	n, m = len(matrix), len(matrix[0])
	//同理裁剪列宽
	startM, endM := 0, m
	for i := 0; i < m; i++ {
		if matrix[n-1][i] < target {
			startM = i+1
		}
		if matrix[0][i] > target {
			endM = i
			break
		}
	}
	if startM >= endM {
		return false
	}
	//裁剪大于部分后的数组
	for i := 0; i < len(matrix); i++ {
		matrix[i] = matrix[i][startM:endM]
	}
	//当裁剪后，如果数字存在数组中，那么一定在左下角或右上角，因为其他的被裁切掉了
	if matrix[0][len(matrix[0])-1] == target || matrix[len(matrix)-1][0] == target {
		return true
	}
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	return my(matrix, target)
}
