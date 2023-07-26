/*
 * @Author: shengchao
 * @Date: 2023-07-24 14:19:08
 * @LastEditors: chaosheng sean.sheng@live.com
 * @LastEditTime: 2023-07-26 23:05:20
 * @Description: 基于map + 有序链表实现优先队列
 * @FilePath: /algo_proj/queue/priority_queue_v1.go
 */
package queue

import (
	"algo_proj/api"
	"algo_proj/list"
	"algo_proj/utils"
	"sync"
)

type Key interface {
	~int | ~string | ~float32
}

// map + slice 实现优先级队列
type PriorityQueueV1[K Key, T interface{}] struct {
	size                int
	descSort            bool               // 优先级是否以降序存放
	mutex               sync.Mutex         // 防止并发操作导致数据不一致
	dataCompareFunc     api.CompareFunc[T] // 真实数据比较函数[必须设置]
	priorityCompareFunc api.CompareFunc[K] // 优先级数据比较函数[必须设置]
	queue               map[K]api.List[T]  // <k,v>=<优先级,数据本身>
	priorities          api.List[K]        // 按降序/升序存放优先级
}

func NewPriorityQueueV1[K Key, T interface{}](desc bool, cmpPriorityFunc api.CompareFunc[K], cmpDataFunc api.CompareFunc[T]) api.PriorityQueue[K, T] {
	opt := utils.Asc
	if desc {
		opt = utils.Desc
	}

	return &PriorityQueueV1[K, T]{
		size:                0,
		descSort:            desc,
		dataCompareFunc:     cmpDataFunc,
		priorityCompareFunc: cmpPriorityFunc,
		priorities:          list.NewSinglyLinkedList[K](opt, cmpPriorityFunc),
	}
}

func (pq *PriorityQueueV1[K, T]) EnQueue(priority K, val T) bool {
	pq.mutex.Lock()
	defer pq.mutex.Unlock()

	v, err := pq.queue[priority]
	if !err { //map中暂时没有该优先级记录
		v = list.NewSinglyLinkedList[T](utils.No, pq.dataCompareFunc) // 创建无序链表
		pq.queue[priority] = v                                        // 优先级与数据链表关联
		pq.priorities.Insert(priority)                                // 按序保存优先级数据
	}

	v.Insert(val) // 将真实数据保存在链表中
	return true
}

func (pq *PriorityQueueV1[K, T]) DeQueue(output *T) bool {
	if output == nil {
		return false
	}

	pq.mutex.Lock()
	defer pq.mutex.Unlock()

	if pq.priorities.Size() <= 0 {
		return false
	}

	priority, err := pq.priorities.Head()
	if err != nil {
		return false
	}

	v, success := pq.queue[priority]
	if !success {
		return success
	}

	*output, err = v.Head()
	if err != nil {
		return false
	}

	v.Delete(*output)
	return true
}

func (pq *PriorityQueueV1[K, T]) Size() int {
	pq.mutex.Lock()
	defer pq.mutex.Unlock()
	return pq.size
}
