package isMatch

import (
	"fmt"
)

//用状态机尝试解决此问题，却导致越陷越深，各种细节和分支处理起来太繁琐，此方法弃用(未完成)
func my(s string, p string) bool {
	//先对p进行分组
	type Reg struct {
		reg string //正则数据
		typ int    //正则类型
	}
	type pGroup struct {
		index int   //分组偏移
		group []Reg //分组数据
		state int   //状态机标识，0是字符，1是. 2是*
	}
	pg := pGroup{
		index: 0,
		group: []Reg{{reg: "", typ: 0}},
		state: -1,
	}
	for _, c := range p {
		switch c {
		case '.': //不论点号前面是什么都重新分一个组，除非以点号开头
			if pg.state == -1 {
				pg.group[pg.index].reg += string(c)
				pg.group[pg.index].typ = 1
			} else {
				pg.group[pg.index].typ = pg.state //前面分组
				pg.group = append(pg.group, Reg{reg: string(c), typ: 1})
				pg.index++
			}
			pg.state = 1
		case '*':
			//星号一定是和前面字符一组，但是两个星号重复的时候可以省略后面的星号简化逻辑
			if pg.state == 2 {
				continue
			}
			//如果前面分组不止有一个字符，需要将前面字符拆分
			if len(pg.group[pg.index].reg) != 1 {
				pg.group[pg.index].typ = pg.state                                               //前面分组
				cc := pg.group[pg.index].reg[len(pg.group[pg.index].reg)-1:]                    //取出最后一个字符
				pg.group[pg.index].reg = pg.group[pg.index].reg[:len(pg.group[pg.index].reg)-1] //前面分组删掉最后一个字符

				pg.group = append(pg.group, Reg{reg: cc + string(c), typ: 2}) //新开分组
				pg.state = 2
				pg.index++
			} else {
				pg.group[pg.index].reg += string(c)
				pg.state = 2
				pg.group[pg.index].typ = 2
			}
		default:
			//如果字符前面是点号或者星号，则需要重新分组，并标记前面分组类型
			if pg.state == 1 || pg.state == 2 {
				pg.group[pg.index].typ = pg.state
				pg.group = append(pg.group, Reg{reg: string(c), typ: 0})
				pg.index++
			} else {
				pg.group[pg.index].reg += string(c)
				pg.group[pg.index].typ = 0
			}
			pg.state = 0
		}
	}
	fmt.Printf("%+v\n", pg)

	//用分组的正则去匹配字符串
	for i, reg := range pg.group {
		switch reg.typ {
		case 0: //纯字符，直接对比，如果相同减去匹配的字符
			if len(s) < len(reg.reg) {
				return false
			}
			if s[:len(reg.reg)] != reg.reg {
				return false
			}
			s = s[len(reg.reg):] //截断已匹配字符
		case 1: //带有点号的规则,任意匹配一个字符，注意上面将点号单独分组了
			if len(s) == 0 { //任意字符不包括无字符
				return false
			}
			s = s[1:] //截断已匹配字符
		case 2: //带有星号的规则
			if reg.reg == ".*" { //任意匹配多个字符，此时需要关心后面还有没有规则，如果有规则.*就只能截止到规则结束
				if i == len(pg.group)-1 { //后面没有规则了
					return true
				}
				scopy := s
				for j, reg2 := range pg.group[i:] {
					if reg2.typ == 1 { //点号必须匹配字符
						if len(s) == 0 { //任意字符不包括无字符
							return false
						}
						scopy = scopy[1:]
					}
					if reg2.typ == 0 { //字符一旦出现立即截止，如果此时j大于0表示有j+1个点号匹配规则
						if len(scopy) < len(reg2.reg) {
							return false
						}
						if scopy[:len(reg2.reg)] != reg2.reg {
							return false
						}
						s = s[len(s)-len(scopy)+j+1:] //s新串等于剩下字符向前推进j+1个点号
						break
					}
				}
			} else {
				//星号可能有多个，但只能有一个前缀，无限匹配这个前缀字符
				var num int
				for _, ch := range s {
					if string(ch) == reg.reg[:1] { //匹配到
						num++
					} else {
						break//一旦未匹配到需要跳出，避免匹配渗透到后面字符
					}
				}
				//TODO:如果出现a*a的情况需要保留后面这个a
				s = s[num:] //剔除匹配字符
			}
		default:
			panic(reg)
		}
	}
	if len(s) > 0 {
		return false
	}
	return true
}
