/*
 * @Author: shengchao
 * @Date: 2023-07-26 14:49:20
 * @LastEditors: your name
 * @LastEditTime: 2023-07-26 18:17:00
 * @Description: --- 链表的通用接口 ---
 * @FilePath: /algo_proj/api/list.go
 */
package api

// 0: 相等
// 1: 大于
// -1: 小于
type CompareFunc[T interface{}] func(T, T) int

type List[T interface{}] interface {
	Size() int
	Head() (T, error)
	Find(val T) bool
	Insert(val T) bool
	Delete(val T) bool
}
