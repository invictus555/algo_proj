package adaptor_containers

import "algo_proj/sequential_containers"

// stack可用适配器:list、vector、deque,说明它们有相同的接口
type Adaptor4Stack[D interface{}] interface {
	sequential_containers.List[D] | sequential_containers.Vector[D] | sequential_containers.Deque[D]
}

// 嵌套范型
type Stack[T interface{}, K Adaptor4Stack[T]] struct {
}
