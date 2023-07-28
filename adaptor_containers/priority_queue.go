package adaptor_containers

import "algo_proj/sequential_containers"

// 优先队列的可用适配器:deque、vector,说明它们有相同的接口
type Adaptor4PriorityQueue[D interface{}] interface {
	sequential_containers.Deque[D] | sequential_containers.Vector[D]
}

type PriorityQueue[T interface{}, K Adaptor4PriorityQueue[T]] struct {
	size      int
	container K
}

func (pq *PriorityQueue[T, K]) Top(val *T) bool {
	return true
}

func (pq *PriorityQueue[T, K]) Push(val T) bool {
	return true
}

func (pq *PriorityQueue[T, K]) Pop(val *T) bool {
	return true
}

func (pq *PriorityQueue[T, K]) Size() int {
	return 0
}

func (pq *PriorityQueue[T, K]) Empty() bool {
	return pq.Size() == 0
}
