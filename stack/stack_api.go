/*
 * @Author: shengchao
 * @Date: 2023-07-26 17:28:39
 * @LastEditors: your name
 * @LastEditTime: 2023-07-26 17:33:20
 * @Description:
 * @FilePath: /algo_proj/stack/stack_api.go
 */

package stack

type Stack[Type interface{}] interface {
	Size() int
	IsEmpty() bool
	Pop() (Type, error)
	Top() (Type, error)
	Push(val Type) bool
}
