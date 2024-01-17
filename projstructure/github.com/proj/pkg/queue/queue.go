package queue

import "fmt"

// Queue represents a basic queue data structure.
type Queue []int

// Enqueue adds an element to the end of the queue.
func (q *Queue) Enqueue(value int) {
	*q = append(*q, value)
}

// Dequeue removes and returns the element from the front of the queue.
func (q *Queue) Dequeue() int {
	if len(*q) == 0 {
		return 0 // or any other appropriate default value
	}
	value := (*q)[0]
	*q = (*q)[1:]
	return value
}

// IsEmpty returns true if the queue is empty.
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

// Peek returns the element at the front of the queue without removing it.
func (q *Queue) Peek() int {
	if len(*q) == 0 {
		return 0 // or any other appropriate default value
	}
	return (*q)[0]
}

func main() {
	queue := make(Queue, 0)

	// Enqueue elements to the queue
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	// Peek at the element at the front of the queue
	fmt.Println("Peeked value:", queue.Peek())

	// Dequeue elements from the queue
	for !queue.IsEmpty() {
		value := queue.Dequeue()
		fmt.Printf("Dequeued value: %d\n", value)
	}
}
