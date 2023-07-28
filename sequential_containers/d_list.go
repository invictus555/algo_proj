package sequential_containers

/* d_list stands for double-linked list */
type DListNode[T interface{}] struct {
	val  T
	next *DListNode[T]
	prev *DListNode[T]
}

func NewDListNode[T interface{}](val T) *DListNode[T] {
	return &DListNode[T]{
		val:  val,
		next: nil,
		prev: nil,
	}
}

func NewDefaultDListNode[T interface{}]() *DListNode[T] {
	return &DListNode[T]{
		next: nil,
		prev: nil,
	}
}

// 基于双链表的无序列表(元素不连续存放)
type DList[T interface{}] struct {
	size   int           // 元素个数
	header *DListNode[T] // 指向双向链表的头部
	tailer *DListNode[T] // 指向双向链表的尾部
}

func (list *DList[T]) Size() int {
	return list.size
}

func (list *DList[T]) Empty() bool {
	return list.Size() == 0
}

func (list *DList[T]) Insert(index int, val T) bool {
	return true
}

func (list *DList[T]) PushFront(val T) bool {
	return true
}

func (list *DList[T]) PopFront(val *T) bool {
	return true
}

func (list *DList[T]) Front(val *T) bool {
	return true
}

func (list *DList[T]) Back(val *T) bool {
	return true
}

func (list *DList[T]) PushBack(val T) bool {
	return true
}

func (list *DList[T]) PopBack(val *T) bool {
	return true
}
