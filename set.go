package main

import "fmt"

// Set represents a basic set data structure.
type Set map[int]struct{}

// Add adds an element to the set.
func (s Set) Add(value int) {
	s[value] = struct{}{}
}

// Remove removes an element from the set.
func (s Set) Remove(value int) {
	delete(s, value)
}

// Contains checks if the set contains a specific element.
func (s Set) Contains(value int) bool {
	_, exists := s[value]
	return exists
}

// Size returns the size of the set.
func (s Set) Size() int {
	return len(s)
}

// Elements returns a slice containing all elements of the set.
func (s Set) Elements() []int {
	elements := make([]int, 0, len(s))
	for element := range s {
		elements = append(elements, element)
	}
	return elements
}

func main() {
	mySet := make(Set)

	// Add elements to the set
	mySet.Add(1)
	mySet.Add(2)
	mySet.Add(3)
	mySet.Add(3)

	// Check if the set contains a specific element
	fmt.Println("Contains 2:", mySet.Contains(2))
	fmt.Println("Contains 4:", mySet.Contains(4))

	// Print the elements of the set
	fmt.Println("Elements:", mySet.Elements())

	// Remove an element from the set
	mySet.Remove(2)

	// Print the elements of the set after removal
	fmt.Println("Elements after removal:", mySet.Elements())
}
