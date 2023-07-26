/*
 * @Author: shengchao
 * @Date: 2023-06-19 11:00:49
 * @LastEditors: your name
 * @LastEditTime: 2023-07-26 18:17:48
 * @Description:--- 范型语法定义单链表 ---
 * @FilePath: /algo_proj/list/singly_linked_list_impl.go
 */
package list

import (
	"algo_proj/api"
	"algo_proj/utils"
)

// 带有头节点的单链表
type SinglyLinkedList[T interface{}] struct {
	size     int
	sortType utils.SortType
	head     *SinglyNode[T]
	cmpFunc  api.CompareFunc[T]
}

func NewSinglyLinkedList[T interface{}](sortType utils.SortType, cmpFunc api.CompareFunc[T]) *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{
		size:     0,
		sortType: sortType,
		cmpFunc:  cmpFunc,
		head:     NewDefaultSinglyNode[T](),
	}
}

func (list *SinglyLinkedList[T]) Insert(val T) bool {
	node := NewSinglyNode(val)
	if node == nil {
		return false
	}

	head := list.head
	targetPosition := list.head
	if list.cmpFunc != nil && list.sortType != utils.No { // 以升序/降序插入
		cmpfunc := list.cmpFunc
		curPointer := head.next
		for ; curPointer != nil; curPointer = curPointer.next {
			// 升序插入需找最近比val小的位置
			if list.sortType == utils.Asc && cmpfunc(curPointer.val, val) <= 0 { // next.val <= val
				targetPosition = curPointer // 记录最后一个次匹配结果所在的指针
				continue
			}
			// 降序插入需找最近比val大的位置
			if list.sortType == utils.Desc && cmpfunc(curPointer.val, val) >= 0 { // next.val >= val
				targetPosition = curPointer // 记录最后一个次匹配结果所在的指针
				continue
			}
			break
		}
	}

	node.next = targetPosition.next
	targetPosition.next = node
	list.size = list.size + 1
	return true
}

func (list *SinglyLinkedList[T]) Delete(val T) bool {
	if list.head.next == nil {
		return false
	}

	var found = false
	head := list.head
	curPointer := head.next
	prePointer := list.head
	for ; curPointer != nil; curPointer = curPointer.next {
		if list.cmpFunc(curPointer.val, val) == 0 {
			found = true
			break
		}
		prePointer = curPointer // 记录上一个位置
	}

	if !found {
		return false
	}

	prePointer.next = curPointer.next
	list.size = list.size - 1
	return true
}

func (list *SinglyLinkedList[T]) Find(val T) bool {
	if list.head.next == nil {
		return false
	}

	pointer := list.head.next
	for {
		if pointer == nil {
			break
		}
		if list.cmpFunc(val, pointer.val) == 0 {
			break
		}
		pointer = pointer.next
	}

	return pointer != nil
}

func (list *SinglyLinkedList[T]) Head() (T, error) {
	if list.head.next == nil {
		return list.head.val, utils.ErrListEmpty
	}
	return list.head.next.val, nil
}

func (list *SinglyLinkedList[T]) Size() int {
	return list.size
}
