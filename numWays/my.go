package numWays

func my(n int) int {
	if n < 2 {
		return 1
	}
	var bp = make([]int, n+1)
	bp[0], bp[1] = 1, 1
	for i := 2; i <= n; i++ {
		bp[i] = (bp[i-1] + bp[i-2])%(1e9 + 7)//提前取模，防止溢出
	}
	return bp[n]
}
