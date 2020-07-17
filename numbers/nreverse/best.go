package nreverse

func best(x int) int {
	y := int32(x) //64位机器上int默认是64
	var ans int32
	for y != 0 {
		if ans*10/10 != ans { //溢出
			ans = 0
			break
		}
		ans = ans*10 + y%10
		y = y / 10
	}
	return int(ans)
}
