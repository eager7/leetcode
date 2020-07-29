package findNumberIn2DArray

import (
	"fmt"
	"reflect"
	"testing"
)

func TestExample(t *testing.T) {
	type Te struct {
		In  [][]int
		In2 int
		Out bool
	}
	array1 := [][]int{
		{1, 2, 3, 10, 18},
		{4, 5, 6, 13, 21},
		{7, 8, 9, 14, 23},
		{11, 12, 16, 17, 26},
		{15, 19, 22, 24, 30},
	}
	array2 := [][]int{
		{1, 1},
	}
	array3 := [][]int{
		{5, 6, 9},
		{9, 10, 11},
		{11, 14, 18},
	}
	array4 := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	array5 := [][]int{
		{-5},
	}
	array6 := [][]int{
		{1, 3, 5},
	}
	var tList = []Te{
		{In: array6, In2: 2, Out: false},
		{In: array6, In2: 3, Out: true},
		{In: array5, In2: -10, Out: false},
		{In: array4, In2: 5, Out: true},
		{In: array2, In2: 1, Out: true},
		{In: array3, In2: 9, Out: true},
		{In: array1, In2: 5, Out: true},
	}
	for _, te := range tList {
		if out := best(te.In, te.In2); out != te.Out {
			t.Fatal(te, out)
		}
	}
}
func TestArray(t *testing.T) {
	nn := []int{1, 2, 3, 4}
	nn2 := [...]int{1, 2, 3, 4}
	fmt.Printf("%T,%T,%T\n", nn, nn2, nn2[:])
	fmt.Println(reflect.TypeOf(nn), reflect.TypeOf(nn2))
	fmt.Println(nn[:4])
	fmt.Println(nn[0:4])
	fmt.Println(nn[2:4])
	fmt.Println(nn[3:4])
}

func TestName(t *testing.T) {
	n := 10
	for i := 0; i < n; i++ {
		n = 5
		fmt.Println(i)
	}
}
