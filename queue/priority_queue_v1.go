/*
 * @Author: shengchao
 * @Date: 2023-07-24 14:19:08
 * @LastEditors: your name
 * @LastEditTime: 2023-07-26 17:16:03
 * @Description:
 * @FilePath: /algo_proj/queue/priority_queue_v1.go
 */
package queue

import (
	"algo_proj/list"
	"algo_proj/utils"
	"fmt"
	"sync"
)

// true: 升序
// false: 降序
type cmpFunc func(int, int) bool

// map + slice 实现优先级队列
type PriorityQueueV1 struct {
	cmp        cmpFunc                                // func函数返回值决定升序/降序
	mutex      sync.Mutex                             // 防止并发操作导致数据不一致
	queue      map[int]*list.SinglyLinkedList[string] // <k,v>=<优先级,数据本身>
	priorities *list.SinglyLinkedList[int]            // 按降序/升序存放优先级
}

func NewPriorityQueue(cmp cmpFunc) *PriorityQueueV1 {
	return &PriorityQueueV1{
		cmp:        cmp,
		priorities: list.NewSinglyLinkedList[int](),
	}
}

func (pq *PriorityQueueV1) Push(priority int, value string) error {
	if priority < 0 {
		return utils.ErrInputParameter
	}

	pq.mutex.Lock()
	defer pq.mutex.Unlock()

	val, err := pq.queue[priority]
	if !err { // map中暂时没有该优先级记录
		val = list.NewSinglyLinkedList[string]()
		pq.queue[priority] = val
		err := pq.priorities.Push(priority)
		fmt.Printf("first time to add priority, %+v\n", err)
	}

	return val.Push(value) // 保存真实数据
}

func (pq *PriorityQueueV1) Pop() (string, error) {
	pq.mutex.Lock()
	defer pq.mutex.Unlock()

	if pq.priorities.Size() <= 0 {
		return "", utils.ErrQueueEmpty
	}

	// 先找到最高/最低优先级的标识
	priority, _ := pq.priorities.Peak()
	list, err := pq.queue[priority]
	if !err {
		return "", utils.ErrQueueEmpty
	}

	val, errs := list.Peak()
	if errs != nil {
		return "", errs
	}

	deletePeak := func() {
		_, err := pq.priorities.Pop()
		fmt.Printf("delete peak data, %+v", err)
	}

	if list.Size() <= 0 { // 该优先级队列已空
		deletePeak()               // 删除最高优先级记录
		delete(pq.queue, priority) // 删除该优先级与正式数据的记录
	}
	return val, nil
}
