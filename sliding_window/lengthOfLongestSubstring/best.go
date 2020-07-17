package lengthOfLongestSubstring

import (
	"fmt"
	"strings"
)

/*
滑动窗口的思路:
窗口可以在两个边界移动一开始窗口大小为0
随着数组下标的前进窗口的右侧依次增大
每次查询窗口里的字符，若窗口中有查询的字符
窗口的左侧移动到该字符加一的位置
每次记录窗口的最大程度
重复操作直到数组遍历完成
返回最大窗口长度即可
*/
func best(s string) int {
	var max = 0
	left, right := 0, 0

	sub := s[left:right]
	for _, r := range s {
		if index := strings.IndexRune(sub, r); index != -1 { //sub中有和r重复的字符，移动left，更新sub
			left = left + index + 1 //窗口左移
		}
		right++
		sub = s[left:right] //更新sub
		fmt.Println("sub:", left, right, sub)
		if max < len(sub) {
			max = len(sub) //重新计算最大长度
		}
	}
	return max
}
