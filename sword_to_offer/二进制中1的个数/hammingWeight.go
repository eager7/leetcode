package hammingWeight

import "fmt"

//映射表方案
func hammingWeight1(num uint32) int {
	if num == 0 {
		return 0
	}
	convert := func(s byte) int {
		switch s {
		case '1', '2', '4', '8':
			return 1
		case '3', '5', '6', '9', 'a', 'c':
			return 2
		case '7', 'b', 'd', 'e':
			return 3
		case 'f':
			return 4
		}
		return 0
	}
	ret := 0
	for _, s := range fmt.Sprintf("%x", num) {
		ret += convert(byte(s))
	}
	return ret
}

//移位方案
func hammingWeight(num uint32) int {
	count := uint32(0)
	for num > 0 {
		count += num & uint32(0x01)
		num = num >> 1
	}
	return int(count)
}
