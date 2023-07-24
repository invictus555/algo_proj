package list

import "algo_proj/utils"

type SinglyNode struct {
	val  string
	next *SinglyNode
}

// 带有头节点的单链表
type SinglyLinkedList struct {
	size int
	head *SinglyNode
}

func NewSinglyNode(val string) *SinglyNode {
	return &SinglyNode{
		val:  val,
		next: nil,
	}
}

func NewSinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{
		size: 0,
		head: NewSinglyNode(""),
	}
}

func (list *SinglyLinkedList) Push(val string) error {
	node := NewSinglyNode(val)
	if node == nil {
		return utils.ErrOom
	}
	// 头插法
	node.next = list.head.next
	list.head.next = node
	list.size = list.size + 1
	return nil
}

func (list *SinglyLinkedList) Pop() (string, error) {
	if list.head.next == nil {
		return "", utils.ErrEmpty
	}

	ret := list.head.next
	list.head.next = list.head.next.next
	list.size = list.size - 1
	return ret.val, nil
}

func (list *SinglyLinkedList) Peak() (string, error) {
	if list.head.next == nil {
		return "", utils.ErrEmpty
	}
	return list.head.next.val, nil
}

func (list *SinglyLinkedList) Size() int {
	return list.size
}
