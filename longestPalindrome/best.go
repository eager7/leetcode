package longestPalindrome

//动态规划实现方法
func best(s string) string {
	if len(s) <= 1 {
		return s
	}

	le, start, maxLen := len(s), 0, 1
	dp := make([][]bool, le)

	for r := 0; r < le; r++ {
		dp[r] = make([]bool, le)
		dp[r][r] = true //单字符一定是回文
		for l := 0; l < r; l++ {
			if s[l] == s[r] && (r-l <= 2 || dp[l+1][r-1]) { //满足状态转移条件
				dp[l][r] = true
			} else {
				dp[l][r] = false
			}

			if dp[l][r] {
				curLen := r - l + 1
				if curLen > maxLen {
					maxLen = curLen
					start = l
				}
			}
		}
	}
	return s[start : start+maxLen]
}
