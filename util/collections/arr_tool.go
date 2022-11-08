package collections

import "reflect"

type ArrayList[T comparable] struct {
	data []T
}

func NewArray[T comparable](src T, ret T) ArrayList[T] {
	return ArrayList[T]{
		make([]T, 0),
	}
}
func (a *ArrayList[T]) IsNil() bool {
	return a == nil || a.data == nil
}
func (a *ArrayList[T]) IsEmpty() bool {
	return a.IsNil() || len(a.data) == 0
}
func (a *ArrayList[T]) Add(data T, index int) {
	if index < 0 || index >= len(a.data) {
		a.data = append(a.data, data)
		return
	}
	a.data = append(a.data, *new(T))
	copy(a.data[index+1:], a.data[index:])
	a.data[index] = data
}
func (a *ArrayList[T]) Equals(billi *ArrayList[T]) bool {
	return reflect.DeepEqual(a.data, billi.data)
}
func (a *ArrayList[T]) Contains(src T) bool {
	if a.IsEmpty() {
		return false
	}
	for _, v := range a.data {
		if v == src {
			return true
		}
	}
	return false
}
func (a *ArrayList[T]) Size() int {
	if a != nil && a.data != nil {
		return len(a.data)
	}
	return 0
}
func (a *ArrayList[T]) Get(index int) T {
	if index >= a.Size() || index < 0 {
		return *new(T)
	}
	return a.data[index]
}
func (a *ArrayList[T]) Clone() *ArrayList[T] {
	if a.IsNil() {
		return nil
	}
	ret := &ArrayList[T]{
		data: make([]T, len(a.data)),
	}
	copy(ret.data, a.data)
	return ret
}
