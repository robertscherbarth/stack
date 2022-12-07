/*
Package stack implement a stack with generics

A stack is a linear data structure that is used to store a collection of objects. It is based on Last-In-First-Out (LIFO). A good example of a stack is a pile of plates in a cafeteria. The plates are added to the top of the pile and removed from the top of the pile. The most recently added plate is the first one to be removed.

The LIFO (Last In First Out) principle entails inserting a single element at the end of a stack, i.e., the element inserted at the end of the stack is the first element to appear.

Operations:

`Pop` operations are the insertion of an element into the stack and the deletion of an element from the stack
`Push` operations are the insertion of elements into the stack

*/
package stack

// Stack describes a stack
type Stack[T any] struct {
	sync.Mutex
	stack []T
}

// New returns a Stack
func New[T any]() *Stack[T] {
	return &(Stack[T]{stack: []T{}})
}

// Len returns the length of the stack
func (s *Stack[T]) Len() int {
	return len((*s).stack)
}

// Push method appends an item to the top of the stack
func (s *Stack[T]) Push(item T) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	(*s).stack = append((*s).stack, item)
}

// Pop removes the last item of the stack and return it to its caller
func (s *Stack[T]) Pop() (item T) {
	if len(s.stack) == 0 {
		return
	}
	s.Mutex.Lock()

	item = (*s).stack[len((*s).stack)-1]
	(*s).stack = (*s).stack[:len((*s).stack)-1]

	s.Mutex.Unlock()

	return item
}

// Peek acts very similarly to the `Pop` method but without actually removing an item from the stack
func (s *Stack[T]) Peek() (item T) {
	if len(s.stack) == 0 {
		return
	}
	s.Mutex.Lock()
	item = (*s).stack[len((*s).stack)-1]
	s.Mutex.Unlock()
	return item
}

// GetStack returns stack content
func (s *Stack[T]) GetStack() (stack []T) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	return (*s).stack
}
