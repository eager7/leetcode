package lengthOfLongestSubstring

/*
从左向右遍历字符串，当有两个相同字符时表示不符合题设条件，需要重新计算最大长度
*/
func my(s string) int {
	var mStr = make(map[rune]int)
	var max = 0

	for i, r := range s {
		if index, ok := mStr[r]; ok {
			//小于等于index的全部删掉
			for k, v := range mStr {
				if v <= index {
					delete(mStr, k)
				}
			}
		}
		mStr[r] = i

		if max < len(mStr) {
			max = len(mStr) //重新计算最大长度
		}
	}
	return max
}
