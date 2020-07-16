package uniquePaths

/*
用动态规划的思路解题
bp[m][n]一定是由bp[m][n-1]和bp[m-1][n]而来，因此可以把问题无限分解成一个斐波那契数列
然后就变成了经典的动态规划问题
需要注意初始条件的设置需要把所有单行单列都设置为1
*/
func my(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}
	var bp = make([][]int, m)//初始化数组
	for i := 0; i < m; i++ {
		bp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		bp[i][0] = 1

	}
	for j := 1; j < n; j++ {
		bp[0][j] = 1
	}

	//注意数组溢出问题
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			bp[i][j] = bp[i-1][j] + bp[i][j-1]
		}
	}
	return bp[m-1][n-1]
}
