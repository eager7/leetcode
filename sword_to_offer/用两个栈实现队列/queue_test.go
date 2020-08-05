package queue

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	st := Stack{}
	st.Push(1)
	st.Push(2)
	fmt.Println(st.Pop())
	fmt.Println(st.Pop())
	fmt.Println(st.Pop())
}

func TestQueue(t *testing.T) {
	qu := Constructor()
	if val := qu.DeleteHead(); val != -1 {
		t.Fatal(val)
	}
	qu.AppendTail(5)
	qu.AppendTail(2)
	if val := qu.DeleteHead(); val != 5 {
		t.Fatal(val)
	}
	if val := qu.DeleteHead(); val != 2 {
		t.Fatal(val)
	}
}
