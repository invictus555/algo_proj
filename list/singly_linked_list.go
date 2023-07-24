package list

type SinglyNode struct {
	val  string
	next *SinglyNode
}

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

func (list *SinglyLinkedList) Push(val string) {
	node := NewSinglyNode(val)

	// 头插法
	node.next = list.head.next
	list.head.next = node
	list.size = list.size + 1
}

func (list *SinglyLinkedList) Pop() string {
	if list.head.next == nil {
		panic("empty")
	}

	ret := list.head.next
	list.head.next = list.head.next.next
	list.size = list.size - 1
	return ret.val
}

func (list *SinglyLinkedList) Peak() string {
	if list.head.next == nil {
		panic("empty")
	}
	return list.head.next.val
}

func (list *SinglyLinkedList) Size() int {
	return list.size
}
