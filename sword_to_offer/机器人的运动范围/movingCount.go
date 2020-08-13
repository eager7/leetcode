package movingCount

func movingCount(m int, n int, k int) int {
	visited := make([][]int, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]int, n)
	}
	return dfs(visited, m, n, 0, 0, k)
}

func dfs(visited [][]int, m, n, start, end, k int) int {
	getSum := func(a int) int {
		sum := a % 10
		tmp := a / 10
		for tmp > 0 {
			sum += tmp % 10
			tmp /= 10
		}
		return sum
	}
	if start >= m || start < 0 || end >= n || end < 0 {
		return 0
	}
	if visited[start][end] == 1 || (getSum(start)+getSum(end)) > k {
		return 0
	}
	visited[start][end] = 1
	return 1 +
		dfs(visited, m, n, start+1, end, k) +
		dfs(visited, m, n, start, end+1, k)
}
