package practice

import (
	"fmt"
	"reflect"
	"testing"
)

func TestStructTag(t *testing.T) {
	type TT struct {
		Name string `json:"name"`
		Age  string `json:"age" gorm:"age"`
	}
	var tt TT

	fmt.Println(tt)
	fmt.Println(reflect.TypeOf(tt))

	v := reflect.TypeOf(tt)
	vName, _ := v.FieldByName("Name")
	fmt.Println(vName.Tag)
	fmt.Println(vName.Tag.Get("json"))
}
