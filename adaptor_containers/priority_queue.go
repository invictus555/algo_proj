package adaptor_containers

import "algo_proj/sequential_containers"

// 优先队列的可用适配器:deque、vector
type Adaptor4PriorityQueue[D interface{}] interface {
	sequential_containers.Deque[D] | sequential_containers.Vector[D]
}

// 嵌套范型
type PriorityQueue[T interface{}, K Adaptor4PriorityQueue[T]] struct {
}
