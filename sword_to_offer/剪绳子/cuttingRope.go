package cuttingRope

import "math"

//动态规划解决方案
func cuttingRope(n int) int {
	dp := make([]int, n)
	dp[0] = n
	dp[1] = (n / 2) * (n - n/2)

	max := int(math.Max(float64(dp[0]), float64(dp[1])))
	for i := 2; i < n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
		if max < dp[i] {
			max = dp[i]
		}
	}
	return max
}
