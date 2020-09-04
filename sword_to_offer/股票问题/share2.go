package share

func share2(prices []int) int {
	dpProfit := make([]int, len(prices)) //第i天的综合利润
	dpShare := make([]bool, len(prices)) //第i天是否还持有股票

	dpProfit[0], dpShare[0] = 0, false //第一天买，第一天卖

}
