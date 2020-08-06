package sort

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	type Te struct {
		In  []int
		Out []int
	}
	tList := []Te{
		{In: []int{2, 4, 3, 6, 5, 8, 9, 7, 1}, Out: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{In: []int{2, 4, 3, 3, 4, 1}, Out: []int{1, 2, 3, 3, 4, 4}},
	}
	for _, te := range tList {
		if out := QuickSort(te.In); !reflect.DeepEqual(te.Out, out) {
			t.Fatal(out)
		}
	}
}

func TestPartition(t *testing.T) {
	var list = []int{8, 2, 4, 3, 5, 8, 9, 7, 1, 6}
	QuickSort(list)
	fmt.Println(list)
}

func BenchmarkQuickSort(b *testing.B) {
	list := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		list = append(list, rand.Intn(1000))
	}
	for i := 0; i < b.N; i++ {
		//QuickSort(list)//BenchmarkQuickSort-4   	    1000	   1528469 ns/op
		QuickSort2(list) //BenchmarkQuickSort-4   	    1000	   1313990 ns/op
	}
}

func TestQuickSort(t *testing.T) {
	list := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		list = append(list, rand.Intn(1000))
	}
	start := time.Now()
	QuickSort2(list)
	fmt.Println(time.Since(start))
}