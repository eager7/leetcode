package cuttingRope

import "math"

/*
** 动态规划解决方案
** 设f(n)为乘积，那么我们要找的就是max(f(n))，而f(n)则有很多种可能性，可以将f(n)分成两个数的乘积f(i)*f(n-i)，
** 同理f(i)和f(n-i)都各自有多种可能性，为了保证取到fn的最大值，我们需要取max(f(i)*f(n-i))，这样就把问题变成了
** 一个递归动态规划问题，从上而下寻找一个最优解，代码则从下而上构建这个最优解
 */
func cuttingRope2(n int) int {
	dp := make([]int, n+1)
	switch n {
	case 0, 1: //n大于1的整数，因此0和1都初始化为0
		return 0
	case 2: //n=2则只能剪成两段，1*1=1
		return 1
	case 3:
		return 2
	case 4:
		return 4
	default:
		dp[0], dp[1], dp[2], dp[3], dp[4] = 0, 1, 2, 3, 4 //初始值的计算比较重要，否则后面没法正确输出
		for i := 4; i <= n; i++ {
			for j := 2; j < i; j++ { //除以2是为了节省[1,5][5,1]这种翻转的计算情况
				if dp[i] < dp[j]*dp[i-j] {
					dp[i] = dp[j] * dp[i-j] //筛选出子段的最大值
				}
			}
		}
		return dp[n]
	}
}

func cuttingRope(n int) int {
	switch n {
	case 0, 1: //n大于1的整数，因此0和1都初始化为0
		return 0
	case 2: //n=2则只能剪成两段，1*1=1
		return 1
	case 3:
		return 2
	case 4:
		return 4
	default:
		times := n / 3
		switch n % 3 {
		case 1: //余数为1，表示可以和一个3组成4，而4可以拆成2*2，大于1*3，因此需要重组一个数
			return int(math.Pow(float64(3), float64(times-1))) * 4
		case 2: //余数为2直接乘上2
			return int(math.Pow(float64(3), float64(times))) * 2
		default: //正好整除，拆成3是最好的结果
			return int(math.Pow(float64(3), float64(times)))
		}
	}
}
