package api

type Key interface {
	~int | ~string | ~float32
}

type Queue[T interface{}] interface {
	EnQueue(val T) bool
	DeQueue(output *T) bool
	Front(output *T) bool
	Back(output *T) bool
	Size() int
}

type PriorityQueue[K Key, T interface{}] interface {
	EnQueue(key K, val T) bool
	DeQueue(output *T) bool
	Size() int
}
