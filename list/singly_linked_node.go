/*
 * @Author: shengchao
 * @Date: 2023-07-26 13:26:12
 * @LastEditors: your name
 * @LastEditTime: 2023-07-26 13:42:17
 * @Description:--- 定义单链表节点 ---
 * @FilePath: /algo_proj/list/singly_linked_node.go
 */
package list

type SinglyNode[Type interface{}] struct {
	val  Type
	next *SinglyNode[Type]
}

func NewSinglyNode[Type interface{}](val Type) *SinglyNode[Type] {
	return &SinglyNode[Type]{
		val:  val,
		next: nil,
	}
}

func NewDefaultSinglyNode[Type interface{}]() *SinglyNode[Type] {
	return &SinglyNode[Type]{
		next: nil,
	}
}
