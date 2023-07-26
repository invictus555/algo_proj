/*
 * @Author: shengchao
 * @Date: 2023-07-26 13:26:12
 * @LastEditors: your name
 * @LastEditTime: 2023-07-26 19:30:16
 * @Description:--- 定义单链表节点 ---
 * @FilePath: /algo_proj/list/linkedlist_node.go
 */
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
