package twoSum

import (
	"fmt"
	"runtime"
	"testing"
)

func TestTwoSum(t *testing.T) {
	runtime.GOMAXPROCS(1)
	fmt.Println(twoSum([]int{3, 4, 2}, 6))
	fmt.Println(twoSum([]int{2, 2, 4, 2}, 4))

	fmt.Println(best([]int{3, 4, 2}, 6))
	fmt.Println(best([]int{2, 2, 4, 2}, 4))
}
