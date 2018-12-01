package DS

import (
	"fmt"
)

type ListNode struct {
	Data interface{}
	Next *ListNode
	Prev *ListNode
}

func (list *ListNode) Traverse() int32 {
	var (
		count int32
		temp  *ListNode
	)
	temp = list
	for temp != nil {
		fmt.Println(temp.Data)
		temp = temp.Next
		count++
	}
	return count
}

func (list *ListNode) ListToBst() *Node {

	if list == nil {
		return nil
	}
	length := list.Traverse()
	return convertToBst(list, 0, length-1)
}

func convertToBst(list *ListNode, start, end int32) *Node {

	if start > end {
		return nil
	}
	var (
		mid, count int32
		temp       *ListNode
	)

	mid = (start + end) / 2
	node := &Node{}

	node.LeftNode = convertToBst(list, start, mid-1)

	temp = list
	for count = 0; count < mid; count++ {
		temp = temp.Next
	}

	node.Data = temp.GetInt()
	node.RightNode = convertToBst(list, mid+1, end)

	fmt.Println(node.Data)
	return node
}

func (list *ListNode) GetInt() int32 {
	var (
		val int32
		ok  bool
	)

	if val, ok = list.Data.(int32); ok {
		return val
	}
	return val
}

func (list *ListNode) FindMiddle() *ListNode {
	var (
		ptr1, ptr2 *ListNode
	)

	ptr1 = list
	ptr2 = list

	for ptr2.Next != nil && ptr2.Next.Next != nil {
		ptr1 = ptr1.Next
		ptr2 = ptr2.Next.Next
	}

	return ptr1
}

func (list *ListNode) Add(data interface{}) {
	node := &ListNode{
		Data: data,
	}

	list.Next = node

}

func InitLinkedList() *ListNode {
	return &ListNode{}
}
