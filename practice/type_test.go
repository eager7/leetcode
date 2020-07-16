package practice

import (
	"fmt"
	"reflect"
	"testing"
)

type Int int    //是一个新类型
type Int2 = int //相当于别名

func TestType(t *testing.T) {
	var i2 Int2 = 2
	fmt.Println(reflect.TypeOf(i2)) //int
}
