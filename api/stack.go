/*
 * @Author: shengchao
 * @Date: 2023-07-26 17:28:39
 * @LastEditors: your name
 * @LastEditTime: 2023-07-26 18:17:25
 * @Description:
 * @FilePath: /algo_proj/api/stack.go
 */

package api

type Stack[T interface{}] interface {
	Size() int
	IsEmpty() bool
	Pop() (T, error)
	Top() (T, error)
	Push(val T) bool
}
