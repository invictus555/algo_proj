package sequential_containers

// 定长数组，不可变动
type Array[T interface{}] struct {
	slices   []T
	capacity int
}

func NewArray[T interface{}](capacity int) *Array[T] {
	if capacity <= 0 {
		panic("input parameter <capacity> is illegal")
	}
	return &Array[T]{
		capacity: capacity,
		slices:   make([]T, capacity),
	}
}

func (arr Array[T]) At(index int) T {
	if index < 0 || index >= arr.capacity {
		panic("out of bound")
	}

	return arr.slices[index]
}

func (arr Array[T]) Size() int {
	return arr.capacity
}
