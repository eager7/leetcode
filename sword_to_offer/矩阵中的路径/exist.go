package exist

func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if dfs(board, word, i, j, 0) { //遍历所有字符
				return true
			}
		}
	}
	return false
}

//深度优先遍历
func dfs(board [][]byte, word string, i, j, k int) bool {
	//超出边界值返回false，字符不匹配返回false
	if i >= len(board) || i < 0 || j >= len(board[0]) || j < 0 || board[i][j] != word[k] {
		return false
	}
	if k == len(word)-1 { //已经匹配到最后一个字符，返回true，表示末端匹配成功
		return true
	}
	temp := board[i][j]
	board[i][j] = '/'
	//运行到这里说明k匹配成功，那么再查找当前字符的上下左右四个方向是否能匹配k+1
	res := dfs(board, word, i+1, j, k+1) ||
		dfs(board, word, i-1, j, k+1) ||
		dfs(board, word, i, j-1, k+1) ||
		dfs(board, word, i, j+1, k+1)
	board[i][j] = temp
	return res
}
