package main

type Stack[T any] struct {
	value []T
	index int
}

// Push()
// Pop()
// Peek()
// Size()

func (stack *Stack[T]) Push(v T) {
	stack.value[stack.index] = v
	stack.index++
}

func (stack *Stack[T]) Pop() {
	stack.value = stack.value[:stack.index]
	stack.index--
}

func (stack *Stack[T]) Size() int {
	return stack.index
}

func (stack *Stack[T]) Peek() *Stack[T] {
	return stack
}
