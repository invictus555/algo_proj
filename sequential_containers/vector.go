package sequential_containers

type Vector[T interface{}] struct {
	slices   []T
	capacity int
}

func NewVector[T interface{}](size int) *Vector[T] {
	if size <= 0 {
		panic("input parameter <size> is illegal")
	}

	return &Vector[T]{
		capacity: size,
		slices:   make([]T, size),
	}
}

func (vec *Vector[T]) Insert() T {

}

func (vec *Vector[T]) Empty() bool {

}

func (vec *Vector[T]) Clear() bool {

}

func (vec *Vector[T]) Front() T {

}

func (vec *Vector[T]) Back() T {

}

func (vec *Vector[T]) PushBack() T {

}

func (vec *Vector[T]) PopBack() T {

}

func (vec *Vector[T]) At(index int) T {

}
