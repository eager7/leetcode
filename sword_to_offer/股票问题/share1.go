package share

func share1(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	dp := make([]int, len(prices)) //dp[i]表示截止到i+1天的最大利润
	dp[0] = 0                      //当天买当天卖就是0
	var min int
	if prices[1] > prices[0] { //dp[i]等于前i+1天
		min, dp[1] = prices[0], prices[1]-prices[0]
	} else {
		min, dp[1] = prices[1], 0
	}

	for i := 2; i < len(prices); i++ { //从第三天开始算
		if prices[i]-min > dp[i-1] { //第i天价格减去前面最小值大于前面最大利润，则更新
			dp[i] = prices[i] - min
		} else {
			dp[i] = dp[i-1]
		}
		if prices[i] < min {
			min = prices[i]
		}
	}
	return dp[len(prices)-1]
}
