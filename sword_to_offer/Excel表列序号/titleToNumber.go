package titleToNumber

import (
	"math"
)

func titleToNumber(s string) int {
	if s == "" {
		return 0
	}
	ret := 0
	for i, c := range s {
		ret += int(c-'A'+1) * int(math.Pow(float64(26), float64(len(s)-1-i)))
	}
	return ret
}
