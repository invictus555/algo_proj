/*
 * @Author: shengchao
 * @Date: 2023-06-18 23:19:34
 * @LastEditors: your name
 * @LastEditTime: 2023-07-26 19:57:14
 * @Description: 基于单链表实现栈
 * @FilePath: /algo_proj/stack/linkedlist_stack.go
 */

package stack

import (
	"algo_proj/api"
	"algo_proj/list"
	"algo_proj/utils"
	"sync"
)

// 基于链表的栈
type LinkedListStack[T interface{}] struct {
	size int         // 栈元素数量
	lock sync.Mutex  // 并发安全锁
	list api.List[T] // 单链表
}

func NewLinkedStack[T interface{}]() api.Stack[T] {
	return &LinkedListStack[T]{
		size: 0,
		list: list.NewSinglyLinkedList[T](utils.No, nil),
	}
}

func (stack *LinkedListStack[T]) Push(v T) bool {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	if err := stack.list.Insert(v); !err {
		return false
	}

	stack.size = stack.size + 1
	return true
}

func (stack *LinkedListStack[T]) Pop() (T, error) {
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

func (stack *LinkedListStack[T]) Top() (T, error) {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	v, err := stack.list.Head()
	if err != nil {
		return v, err
	}
	return v, err
}

// 栈大小
func (stack *LinkedListStack[T]) Size() int {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	return stack.size
}

// 栈是否为空
func (stack *LinkedListStack[T]) IsEmpty() bool {
	return stack.Size() == 0
}
