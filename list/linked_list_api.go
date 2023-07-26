/*
 * @Author: shengchao
 * @Date: 2023-07-26 14:49:20
 * @LastEditors: your name
 * @LastEditTime: 2023-07-26 17:20:03
 * @Description: --- 链表的通用接口 ---
 * @FilePath: /algo_proj/list/linked_list_api.go
 */
package list

// 0: 相等
// 1: 大于
// -1: 小于
type CompareFunc[Type interface{}] func(Type, Type) int

type LinkedList[Type interface{}] interface {
	Size() int
	Head() (Type, error)
	Find(val Type) bool
	Insert(val Type) bool
	Delete(val Type) bool
}
