package main

import "fmt"

// Stack represents a basic stack data structure.
type Stack []int

// Push adds an element to the top of the stack.
func (s *Stack) Push(value int) {
	*s = append(*s, value)
}

// Pop removes and returns the element from the top of the stack.
func (s *Stack) Pop() int {
	if len(*s) == 0 {
		return 0 // or any other appropriate default value
	}
	index := len(*s) - 1
	value := (*s)[index]
	*s = (*s)[:index]
	return value
}

// IsEmpty returns true if the stack is empty.
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Peek returns the element at the top of the stack without removing it.
func (s *Stack) Peek() int {
	if len(*s) == 0 {
		return 0 // or any other appropriate default value
	}
	return (*s)[len(*s)-1]
}

func main() {
	stack := make(Stack, 0)

	// Push elements onto the stack
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	// Peek at the element at the top of the stack
	fmt.Println("Peeked value:", stack.Peek())

	// Pop elements from the stack
	for !stack.IsEmpty() {
		value := stack.Pop()
		fmt.Printf("Popped value: %d\n", value)
	}
}
