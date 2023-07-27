package sequential_containers

import "fmt"

// 可变长数组(元素顺序存放)，复杂度如下:
// 1.查找复杂度:O(1)
// 2.插入/删除复杂度:O(n)
type Vector[T interface{}] struct {
	size     int
	slices   []T
	capacity int
}

/**
 * @description:创建一个Vector实例
 * @param {size}:限定vector初始大小
 * @return {*Vector}: vector实例指针
 */
func NewVector[T interface{}](capacity int) *Vector[T] {
	if capacity <= 0 {
		panic("input parameter <size> is illegal")
	}

	return &Vector[T]{
		size:     0,
		capacity: capacity,
		slices:   make([]T, capacity),
	}
}

/**
 * @description: 创建一个默认容量的vector
 * @return {*Vector}: vector实例指针
 */
func NewDefaultVector[T interface{}]() *Vector[T] {
	return NewVector[T](16)
}

/**
 * @description: 在指定位置插入一个元素
 * @param {int} pos:下标位置[0,...,n-1]
 * @param {T} val: 被插入的数据元素
 * @return {*} successfully if true, otherwise failed
 */
func (vec *Vector[T]) Insert(pos int, val T) bool {
	if pos < 0 {
		panic("input parameter <pos> is illegal")
	}

	if pos >= vec.capacity { // 位置超过容量，须扩容
		ok := enlargeSlice[T](vec)
		fmt.Println("enlarge result:", ok)
	}

	vec.size = vec.size + 1
	vec.slices[pos] = val
	return true
}

func (vec *Vector[T]) Empty() bool {
	return vec.size == 0
}

func (vec *Vector[T]) Clear() bool {

}

func (vec *Vector[T]) Front() T {

}

func (vec *Vector[T]) Back() T {

}

func (vec *Vector[T]) PushBack(val T) bool {
	vec.slices = append(vec.slices, val)
	vec.capacity = vec.capacity + 1
	return true
}

func (vec *Vector[T]) PopBack() T {

}

func (vec *Vector[T]) At(index int) T {
	if index < 0 {
		panic("input parameter <index> is illegal")
	}

	return vec.slices[index]
}

/**
 * @description: slice扩容，遵循准则: 2倍，1.25倍
 * @return {bool} successfully if true, or failed
 */
func enlargeSlice[T interface{}](vec *Vector[T]) bool {
	if vec == nil {
		return false
	}

	capacity := vec.capacity
	if vec.capacity < 1024 {
		vec.capacity = 2 * vec.capacity
	} else if vec.capacity >= 1024 {
		capability := float32(vec.capacity)
		vec.capacity = int(1.25 * capability)
	}

	slices := make([]T, vec.capacity)
	copy(slices[0:capacity-1], vec.slices[0:capacity-1])
	vec.slices = slices // 注意扩容后内存地址的改变
	return true
}
