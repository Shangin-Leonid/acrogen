package cont /* Containers */

// Container that implements stack (LIFO data structure). Based on array.
type Stack[V any] struct {
	data []V
}

// Creates empty stack with 'capacity'.
func NewStack[V any](capacity int) Stack {
	return Stack[V]{data: make([]V, 0, capacity)}
}

// Returns size of stack (amount of elements).
func (s Stack[V]) Size() int {
	return len(s.data)
}

// Returns if stack is empty.
func (s Stack[V]) IsEmpty() bool {
	return s.Size() == 0
}

// Returns the capacity of stack.
func (s Stack[V]) Capacity() int {
	return cap(s.data)
}

// Returns top element (last added element).
// Panics if stack is empty.
func (s Stack[V]) Top() V {
	return s.data[len(s.data)-1]
}

// Returns bottom element (first added element).
// Panics if stack is empty.
func (s Stack[V]) Bottom() V {
	return s.data[0]
}

// Deletes top (last added) element.
// The stack's size decreases by 1.
// Do nothing if stack is already empty.
func (s *Stack[V]) Pop() {
	if !s.IsEmpty() {
		s.data = s.data[:len(s.data)-1]return
	}
}

// Adds a new element to the stack's top.
// The stack's size increases by 1.
func (s *Stack[V]) Push(value V) {
	s.data = append(s.data, value)
}
