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

func TestNewParam(t *testing.T) {
	a := 1
	fmt.Printf("%p\n", &a)
	a, b := 2, 3 //变量覆盖，b是新变量，因此可以用:=赋值
	fmt.Printf("%p,%p\n", &a, &b)
	{
		a, b := 2, 3 //定义新变量，不在一个作用域
		fmt.Printf("%p,%p\n", &a, &b)
	}
}

func TestConst(t *testing.T) {
	const (
		s  = ""
		s1          //正确，和上一个常量相等
		ss = len(s) //正确，s是常量，编译期间可以计算出len
		//sss = len([]int{1, 2}) //错误，切片在编译期间无法计算长度
		_  = iota //通常用于初始化第一个值
		s5 = 5
		s6 //等于6
	)
}
func TestComplex(t *testing.T) {
	fmt.Println(complex(1, 2))
}

func TestInterface(t *testing.T) {
	type User struct {
		name string
	}

	u := User{name: "test"}
	var vi, pi interface{} = u, &u //这里会发生值拷贝，vi得到完全复制体，pi得到指针复制体

	fmt.Println(vi.(User).name)  //test
	fmt.Println(pi.(*User).name) //test

	//vi.(User).name = "vi name"  //ERR:强制转换返回的是临时对象，内部有拷贝，不允许赋值
	pi.(*User).name = "pi name" //返回的是临时指针对象，可以赋值

	fmt.Println(vi.(User).name)  //test，还是拷贝对象
	fmt.Println(pi.(*User).name) //pi name，pi对象是拷贝指针对象，但是指向的具体数据还是u
}
