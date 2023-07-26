/*
 * @Author: shengchao
 * @Date: 2023-07-26 17:36:14
 * @LastEditors: your name
 * @LastEditTime: 2023-07-26 19:20:37
 * @Description:
 * @FilePath: /algo_proj/test/stack_test.go
 */

package test

import (
	"algo_proj/stack"
	"fmt"
	"testing"
)

func TestArrayStack(_ *testing.T) {
	stack := stack.NewArrayStack[string]()
	stack.Push("cat")
	stack.Push("dog")
	stack.Push("hen")
	fmt.Println("size:", stack.Size())
	fmt.Printf("pop: %v, %+v\n", stack.Pop())
	fmt.Printf("pop: %v, %+v\n", stack.Pop())
	fmt.Println("size:", stack.Size())
	stack.Push("drag")
	fmt.Printf("pop: %v, %+v\n", stack.Pop())
}

func TestLinkedStack(_ *testing.T) {
	stack := stack.NewLinkedStack[string]()
	stack.Push("cat")
	stack.Push("dog")
	stack.Push("hen")
	fmt.Println("size:", stack.Size())
	fmt.Printf("pop: %v, %+v\n", stack.Pop())
	fmt.Printf("pop: %v, %+v\n", stack.Pop())
	fmt.Println("size:", stack.Size())
	stack.Push("drag")
	fmt.Printf("pop: %v, %+v\n", stack.Pop())
}
