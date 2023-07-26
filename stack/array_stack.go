/*
 * @Author: shengchao
 * @Date: 2023-06-18 23:19:34
 * @LastEditors: your name
 * @LastEditTime: 2023-07-26 17:35:53
 * @Description:
 * @FilePath: /algo_proj/stack/array_stack.go
 */
package stack

import (
	"sync"
)

// 基于数组的栈
type ArrayStack struct {
	size  int        // 栈的元素数量
	array []string   // 底层切片
	lock  sync.Mutex // 为了并发安全使用的锁
}

// 入栈时直接把元素放在数组的最后面，然后元素数量加 1;
// 性能损耗主要花在切片追加元素上，切片如果容量不够会自动扩容，
// 底层损耗的复杂度我们这里不计，所以时间复杂度为O(1)
func (stack *ArrayStack) Push(v string) bool {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	// 放入切片中，后进的元素放在数组最后面
	stack.array = append(stack.array, v)

	// 栈中元素数量+1
	stack.size = stack.size + 1
	return true
}

// Pop 出栈
func (stack *ArrayStack) Pop() string {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	// 栈中元素已空
	if stack.size == 0 {
		panic("empty")
	}

	// 获取栈顶元素
	v := stack.array[stack.size-1]
	stack.array = stack.array[0 : stack.size-1]

	stack.size = stack.size - 1
	return v
}

// Peak 获取栈顶元素，不出栈
func (stack *ArrayStack) Peak() string {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	// 栈中元素已空
	if stack.size == 0 {
		panic("empty")
	}

	// 栈顶元素值
	v := stack.array[stack.size-1]
	return v
}

// 栈大小
func (stack *ArrayStack) Size() int {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	return stack.size
}

// 栈是否为空
func (stack *ArrayStack) IsEmpty() bool {
	return stack.Size() == 0
}
