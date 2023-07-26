/*
 * @Author: shengchao
 * @Date: 2023-07-26 17:36:14
 * @LastEditors: your name
 * @LastEditTime: 2023-07-26 17:37:18
 * @Description:
 * @FilePath: /algo_proj/test/array_stack_test.go
 */

package test

import (
	"fmt"
	"testing"
)

func TestStack(_ *testing.T) {
	arrayStack := new(ArrayStack)
	arrayStack.Push("cat")
	arrayStack.Push("dog")
	arrayStack.Push("hen")
	fmt.Println("size:", arrayStack.Size())
	fmt.Println("pop:", arrayStack.Pop())
	fmt.Println("pop:", arrayStack.Pop())
	fmt.Println("size:", arrayStack.Size())
	arrayStack.Push("drag")
	fmt.Println("pop:", arrayStack.Pop())
}
