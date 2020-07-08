package myAtoi

//用状态机避免臃肿代码
func best(str string) int {
	type state struct {
		flag int //正负号
		stat int //状态，0-初始化，1-符号处理，2-数字处理，3-结束
		num  int32
	}

	var s state
	s.flag = 1
loop:
	for _, c := range str {
		switch s.stat {
		case 0:
			if c == ' ' { //跳过空格
				continue
			}
			if c == '-' {
				s.flag = -1
			} else if c == '+' {
				s.flag = 1
			} else if c < '0' || c > '9' {
				return 0
			} else {
				s.num = c - '0'
			}
			s.stat = 1
		case 1:
			if c < '0' || c > '9' {
				return 0
			}
			s.num = s.num*10 + c - '0'
			s.stat = 2
		case 2:
			switch c {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				if s.num*10/10 != s.num || s.num*10+c-'0' < s.num { //溢出
					if s.flag == 1 {
						return 1<<31 - 1
					} else {
						return -1 << 31
					}
				}
				s.num = s.num*10 + c - '0'
			default:
				break loop
			}
		}
	}
	return s.flag * int(s.num)
}
