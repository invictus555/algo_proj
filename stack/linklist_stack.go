package stack

import (
	"algo_proj/list"
	"sync"
)

// 基于链表的栈
type LinkedListStack struct {
	size int                   // 栈元素数量
	lock sync.Mutex            // 并发安全锁
	list list.SinglyLinkedList // 单链表
}

func (stack *LinkedListStack) Push(v string) {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	stack.list.Push(v)
	stack.size = stack.size + 1
}

func (stack *LinkedListStack) Pop() string {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	if stack.size <= 0 {
		panic("empty")
	}

	v := stack.list.Pop()
	stack.size = stack.size - 1
	return v
}

func (stack *LinkedListStack) Peak() string {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	if stack.size <= 0 {
		panic("empty")
	}

	return stack.list.Peak()
}

// 栈大小
func (stack *LinkedListStack) Size() int {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	return stack.size
}

// 栈是否为空
func (stack *LinkedListStack) IsEmpty() bool {
	return stack.Size() == 0
}
