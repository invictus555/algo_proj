package adaptor_containers

import "algo_proj/sequential_containers"

// 优先队列的可用适配器:deque、list
type Adaptor4Queue[D interface{}] interface {
	sequential_containers.List[D] | sequential_containers.Deque[D]
}

// 嵌套范型
type queue[T interface{}, K Adaptor4Queue[T]] struct {
}
