package myPow

//暴力循环破解
func my(x float64, n int) float64 {
	if x == 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	ret := x
	switch n < 0 {
	case false:
		for i := 1; i < n; i++ {
			ret = ret * x
		}
	case true:
		n = -n //取绝对值
		for i := 1; i < n; i++ {
			ret = ret * x
		}
		ret = 1 / ret
	}
	return ret
}

/*
** 采用斐波那契数列形式递归求解，可以节省大量计算资源
** 思路即x的32次方等于x的2次方的平方(4)的平方(8)的平方(16)的平方(32)
 */
func myPow(x float64, n int) float64 {
	pow := func(exp int) float64 {
		switch exp {
		case 0:
			return 1
		case 1:
			return x
		default:
			ret := myPow(x, exp>>1) //使用移位替代除2
			ret = ret * ret         //取平方
			if n&0x01 == 1 {        //n是奇数，则需要额外乘上x
				ret = ret * x
			}
			return ret
		}
	}
	if x == 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	ret := x
	switch n < 0 {
	case false:
		ret = pow(n)
	case true:
		ret = pow(-n)
		ret = 1 / ret
	}
	return ret
}
