package DS

import "fmt"

var keyCount int64

type Stack struct {
	stackTray []interface{}
	top       int32
	capacity  int32
	size      int32
}

func (stk *Stack) Push(val interface{}) {
	if stk.top <= stk.capacity {
		stk.top += 1
		stk.stackTray[stk.top] = val
		stk.size++
	} else {
		fmt.Println("Overflow")
	}
}

func (stk *Stack) Pop() interface{} {
	if stk.top > -1 {
		top := stk.top
		val := stk.stackTray[top]
		stk.stackTray[top] = nil
		stk.top = top - 1
		stk.size--
		return val
	} else {
		fmt.Println("Underflow")
		return 0
	}

}

func (stk *Stack) Read() interface{} {
	return stk.stackTray[stk.top]
}

func (stk *Stack) Size() int32 {
	return stk.size
}

func (stk *Stack) IsEmpty() bool {
	if stk.top == -1 {
		return true
	}
	return false
}

func (stk *Stack) AddAll(s *Stack) {
	for !stk.IsEmpty() {
		s.Push(stk.Pop())
	}
}

func InitStack(capacity int32) *Stack {

	stackImp := Stack{
		capacity:  capacity,
		top:       -1,
		size:      0,
		stackTray: make([]interface{}, capacity),
	}
	return &stackImp
}

func InitHeapStack(capacity int64) *HeapNode {
	arr := make([]int64, capacity)
	key := make([]int64, capacity)
	keyCount = 1
	return &HeapNode{
		Capacity: capacity,
		Arr:      arr,
		Key:      key,
	}
}

func (heap *HeapNode) Push(val int64) {
	if heap.InsertKey(keyCount, val) {
		keyCount++
	} else {
		fmt.Println("Overflow")
	}
}

func (heap *HeapNode) Pop() int64 {
	_, val := heap.DeleteKey()
	if val < 0 {
		fmt.Println("Underflow")
	}
	return val
}

func (heap *HeapNode) GetMaxOrMinStack() int64 {
	return heap.GetMaxOrMin()
}
