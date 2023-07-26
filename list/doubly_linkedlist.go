/*
 * @Author: shengchao
 * @Date: 2023-06-19 11:01:21
 * @LastEditors: your name
 * @LastEditTime: 2023-07-26 19:45:48
 * @Description:
 * @FilePath: /algo_proj/list/doubly_linkedlist.go
 */
package list

import (
	"algo_proj/api"
	"sync"
)

type DoublyLinkedList[T interface{}] struct {
	size   int
	lock   sync.Mutex
	header *DoublyLinkedListNode[T]
}

func NewDoublyLinkedList[T interface{}]() api.List[T] {
	return &DoublyLinkedList[T]{
		size:   0,
		header: NewDoublyLinkedListNode[T](),
	}
}

func (list *DoublyLinkedList[T]) Size() int {
	list.lock.Lock()
	defer list.lock.Unlock()

	return list.size
}

func (list *DoublyLinkedList[T]) Head() (T, error) {
	return list.header.val, nil
}

func (list *DoublyLinkedList[T]) Find(val T) bool {
	return true
}

func (list *DoublyLinkedList[T]) Insert(val T) bool {
	return true
}

func (list *DoublyLinkedList[T]) Delete(val T) bool {
	return true
}
