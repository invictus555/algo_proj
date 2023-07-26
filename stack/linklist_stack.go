package stack

import (
	"algo_proj/list"
	"algo_proj/utils"
	"sync"
)

// 基于链表的栈
type LinkedListStack[Type interface{}] struct {
	size int                          // 栈元素数量
	lock sync.Mutex                   // 并发安全锁
	list *list.SinglyLinkedList[Type] // 单链表
}

func NewLinkedStack[Type interface{}](sortType utils.SortType, cmpfunc list.CompareFunc[Type]) *LinkedListStack[Type] {
	return &LinkedListStack[Type]{
		size: 0,
		list: list.NewSinglyLinkedList[Type](sortType, cmpfunc),
	}
}

func (stack *LinkedListStack[Type]) Push(v Type) bool {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	if err := stack.list.Insert(v); !err {
		return false
	}

	stack.size = stack.size + 1
	return true
}

func (stack *LinkedListStack[Type]) Pop() (Type, error) {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	v, err := stack.list.Head()
	if err != nil {
		return v, err
	}

	stack.list.Delete(v)
	stack.size = stack.size - 1
	return v, nil
}

func (stack *LinkedListStack[Type]) Top() (Type, error) {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	v, err := stack.list.Head()
	if err != nil {
		return v, err
	}
	return v, err
}

// 栈大小
func (stack *LinkedListStack[Type]) Size() int {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	return stack.size
}

// 栈是否为空
func (stack *LinkedListStack[Type]) IsEmpty() bool {
	return stack.Size() == 0
}
