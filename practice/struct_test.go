package practice

import (
	"fmt"
	"testing"
)

func TestComp(t *testing.T) {
	s1, s2 := S1{}, S1{}
	fmt.Println(s1 == s2)

	s3, s4 := S2{}, S2{}
	fmt.Println(s3 == s4)
}

func TestSlice(t *testing.T) {
	var tt = []string{
		0:  "a",
		10: "b",
	}
	fmt.Println(len(tt), tt)
}
