package longestPalindrome

func my(s string) string {
	isPalindrome := func(str string) bool {
		le := len(str)
		if le == 1 {
			return true
		}
		for i := 0; i < le/2; i++ {
			if str[i] != str[le-1-i] {
				return false
			}
		}
		return true
	}
	if isPalindrome(s) {
		return s
	}
	for i := 1; i < len(s); i++ {
		windowLen := len(s) - i
		for j := 0; j <= i; j++ {
			//fmt.Println("i:", i, "j:", j, "r:", j+windowLen, "str:", s[j:j+windowLen])
			if isPalindrome(s[j : j+windowLen]) {
				return s[j : len(s)-(i-j)]
			}
		}
	}
	return ""
}
