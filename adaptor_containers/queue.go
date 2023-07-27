package adaptor_containers

import "algo_proj/sequential_containers"

// 优先队列的可用适配器:deque、list,说明它们有相同的接口
type Adaptor4Queue[D interface{}] interface {
	sequential_containers.List[D] | sequential_containers.Deque[D]
}

// 嵌套范型
type queue[T interface{}, K Adaptor4Queue[T]] struct {
	size      int
	container K
}

func (q *queue[T, K]) Push(val T) bool {
	return true
}

func (q *queue[T, K]) Pop(val *T) bool {
	return true
}

func (q *queue[T, K]) Front(val *T) bool {
	return true
}

func (q *queue[T, K]) Back(val *T) bool {
	return true
}

func (q *queue[T, K]) Size() int {
	return 0
}

func (q *queue[T, K]) Empty() bool {
	return q.Size() == 0
}
