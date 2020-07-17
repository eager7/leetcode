package myAtoi

func my(str string) int {
	if len(str) == 0 {
		return 0
	}
	var flag int32 = 1

prefix:
	for i, s := range str {
		switch s {
		case '+', '-':
			if len(str) == i+1 {
				return 0
			}
			str = str[i+1:]
			if s == '-' {
				flag = -1
			}
			break prefix
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			flag = 1
			str = str[i:]
			break prefix
		case ' ':
			continue
		default:
			return 0
		}
	}

	var result int32
start:
	for _, s := range str {
		switch s {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			if result*10/10 != result || result*10+s-48 < result { //溢出，前者溢出两次，后者溢出一次
				if flag == 1 {
					return 1<<31 - 1
				} else {
					return -1 << 31
				}
			}
			result = result*10 + s - 48
		default:
			break start
		}
	}
	return int(flag * result)
}
