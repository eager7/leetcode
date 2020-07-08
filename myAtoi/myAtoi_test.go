package myAtoi

import (
	"fmt"
	"math"
	"testing"
)

func TestMyAtoi(t *testing.T) {
	type Te struct {
		In  string
		Out int
	}
	fmt.Println(math.MaxInt32, math.MinInt32)
	var list = []Te{
		//{In: "-42", Out: -42},
		//{In:"4193 with words", Out: 4193},
		//{In: "words and 987", Out: 0},
		//{In: "-91283472332", Out: -2147483648},
		//{In: "   -42", Out: -42},
		//{In: "   -", Out: 0},
		//{In: "  0000000000012345678", Out: 12345678},
		//{In: "2147483646", Out: 2147483646},
		//{In: "2147483648", Out: 2147483647},
		{In: "-6147483648", Out: -2147483648},
	}
	for _, te := range list {
		if out := my(te.In); out != te.Out {
			t.Fatal(fmt.Sprintf("out:%d, err:%d", te.Out, out))
		}
		if out := best(te.In); out != te.Out {
			t.Fatal(fmt.Sprintf("out:%d, err:%d", te.Out, out))
		}
	}
}

func TestName(t *testing.T) {
	fmt.Println(math.MaxInt32, math.MinInt32)
	var result = int32(614748364)
	temp := int32(result*10+56-48)
	fmt.Println(temp)//1852516352 溢出两次
	fmt.Println(614748364*10+56-48)//6147483648 被默认转换成64位，不会溢出
}