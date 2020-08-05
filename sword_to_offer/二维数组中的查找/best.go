package findNumberIn2DArray

//官方思路，不需要裁剪数组，直接忽略外围数据即可
func best(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	n, m := len(matrix), len(matrix[0])

	sn, en, sm, em := 0, n-1, 0, m-1
	for sn <= en && sm <= em {
		temp1, temp2 := matrix[en][sm], matrix[sn][em]//取出左上角和右下角数字
		if temp1 < target {
			sm++ //消掉第一行
		} else if temp1 > target {
			en-- //消掉最后一列
		} else {
			return true
		}

		if temp2 < target { //消掉第一列
			sn++
		} else if temp2 > target { //消掉最后一行
			em--
		} else {
			return true
		}
	}
	for i:=sn;i<en;i++{
		for j:=sm;j<em;j++{
			if matrix[i][j]==target{//可以包含查询右下角的元素
				return true
			}
		}
	}
	return false
}
