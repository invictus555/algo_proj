package list

type SinglyLinkedListNode[T interface{}] struct {
	val  T
	next *SinglyLinkedListNode[T]
}

func NewSinglyLinkedListNode[T interface{}](val T) *SinglyLinkedListNode[T] {
	return &SinglyLinkedListNode[T]{
		val:  val,
		next: nil,
	}
}

func NewDefaultSinglyLinkedListNode[T interface{}]() *SinglyLinkedListNode[T] {
	return &SinglyLinkedListNode[T]{
		next: nil,
	}
}

type DoublyLinkedListNode[T interface{}] struct {
	val  T
	next *DoublyLinkedListNode[T]
	prev *DoublyLinkedListNode[T]
}

func NewDoublyLinkedListNode[T interface{}]() *DoublyLinkedListNode[T] {
	return &DoublyLinkedListNode[T]{
		next: nil,
		prev: nil,
	}
}
