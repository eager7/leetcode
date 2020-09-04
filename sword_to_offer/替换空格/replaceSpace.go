package replaceSpace

import "unicode"

func replaceSpace(s string) string {
	var ret string
	for _, v := range s {
		if unicode.IsSpace(v) {
			ret += "%20"
		} else {
			ret += string(v)
		}
	}
	return ret
}
