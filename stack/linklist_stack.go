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

	if err := stack.list.Push(v); err != nil {
		return
	}

	stack.size = stack.size + 1
}

func (stack *LinkedListStack) Pop() (string, error) {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	if stack.size <= 0 {
		panic("empty")
	}

	v, err := stack.list.Pop()
	if err != nil {
		return v, err
	}
	stack.size = stack.size - 1
	return v, nil
}

func (stack *LinkedListStack) Peak() string {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	if stack.size <= 0 {
		panic("empty")
	}

	v, err := stack.list.Peak()
	if err != nil {
		return ""
	}
	return v
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
