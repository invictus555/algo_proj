/*
 * @Author: shengchao
 * @Date: 2023-07-26 13:26:12
 * @LastEditors: your name
 * @LastEditTime: 2023-07-26 18:14:27
 * @Description:--- 定义单链表节点 ---
 * @FilePath: /algo_proj/list/linked_list_node.go
 */
package list

type SinglyNode[T interface{}] struct {
	val  T
	next *SinglyNode[T]
}

func NewSinglyNode[T interface{}](val T) *SinglyNode[T] {
	return &SinglyNode[T]{
		val:  val,
		next: nil,
	}
}

func NewDefaultSinglyNode[T interface{}]() *SinglyNode[T] {
	return &SinglyNode[T]{
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
