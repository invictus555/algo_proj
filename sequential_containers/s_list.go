package sequential_containers

/* s_list stands for single-linked list */

type SListNode[T interface{}] struct {
	val  T
	next *SListNode[T]
}

func NewSListNode[T interface{}](val T) *SListNode[T] {
	return &SListNode[T]{
		val:  val,
		next: nil,
	}
}

func NewDefaultSListNode[T interface{}]() *SListNode[T] {
	return &SListNode[T]{
		next: nil,
	}
}

// 基于单链表的无序列表(元素不连续存放)
// 由于是单链表, 因此list很多特性无法被实现, 例如:back、push_back、pop_back
type SList[T interface{}] struct {
	size   int           // 元素个数
	header *SListNode[T] // 伪头部节点
}

/**
 * @description: 访问表首的元素，不弹出
 * @param {*T} val: 用于携带表首的元素
 * @return {bool}: successfully if true, or failed
 */
func (list *SList[T]) Front(val *T) bool {
	if list.header.next == nil {
		return false
	}

	val = &list.header.val
	return true
}

func (list *SList[T]) Size() int {
	return list.size
}

/**
 * @description: 链表是否为空
 * @return {bool}: yes if true, or no
 */
func (list *SList[T]) Empty() bool {
	return list.Size() == 0
}

/**
 * @description:向单链表某位置的前面插入元素
 * @param {int} index:位置,取值[0,...,n-1]
 * @param {T} val: 元素值
 * @return {bool}: true if successfully, or failed
 * @notice: index下标相当于迭代器位置，有越界的风险
 */
func (list *SList[T]) Insert(index int, val T) bool {
	if index < 0 || list.size <= index {
		panic("index out of boundary")
	}

	curPos := list.header
	for i := index; curPos != nil && i > 0; curPos = curPos.next {
		i--
	}

	node := NewSListNode[T](val)
	node.next = curPos.next
	curPos.next = node
	list.size = list.size + 1
	return true
}

/**
 * @description:弹出链表首的元素
 * @param {*T} val: 指针类型，携带出表首元素
 * @return {bool} successfully if true, or failed
 */
func (list *SList[T]) PopFront(val *T) bool {
	if list.header.next == nil {
		return false
	}

	node := list.header.next
	list.header.next = node.next
	val = &node.val
	return true
}
