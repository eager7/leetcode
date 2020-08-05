package queue

type Stack struct {
	len  int
	data []int
}

func (s *Stack) Push(val int) {
	s.data = append(s.data, val)
}

//先进后出
func (s *Stack) Pop() (end bool, val int) {
	if len(s.data) == 0 {
		return true, -1
	}
	val = s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return false, val
}

type CQueue struct {
	stack1 Stack
	stack2 Stack
}

func Constructor() CQueue {
	return CQueue{
		stack1: Stack{},
		stack2: Stack{},
	}
}

func (this *CQueue) AppendTail(value int) {
	this.stack1.Push(value)
}

//先进先出
func (this *CQueue) DeleteHead() int {
	for {
		end, val := this.stack1.Pop()
		if end {
			break
		}
		this.stack2.Push(val)
	}
	_, resp := this.stack2.Pop()
	for {
		end, val := this.stack2.Pop()
		if end {
			break
		}
		this.stack1.Push(val)
	}
	return resp
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
