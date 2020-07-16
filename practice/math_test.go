package practice

import (
	"fmt"
	"github.com/shopspring/decimal"
	"math"
	"testing"
)

func TestMath(t *testing.T) {
	fmt.Println(math.MaxInt64)
	fmt.Println(math.MaxFloat64)
	mf := decimal.NewFromFloat(math.MaxFloat64)
	fmt.Println(mf.String())
}

func TestFloat(t *testing.T) {
	v := decimal.New(1, 0)
	f := 1.0
	for {
		f++
		v = v.Add(decimal.New(1, 0))
		vv, _ := v.Float64()
		if vv != f {
			t.Fatal(v, f)
		}
		fmt.Println(vv, f)
		if f == math.MaxFloat64 {
			break
		}
	}
	fmt.Println(v, f)
}
