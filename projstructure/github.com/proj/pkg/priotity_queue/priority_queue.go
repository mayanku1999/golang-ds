package priotity_queue

import (
	"errors"
	"fmt"
	"sort"
)

type Item struct {
	value    string
	priority int
}
type PriorityQueue struct {
	Items       []*Item
	HeapTypeMin bool
}

func (pq *PriorityQueue) Enqueue(value string, priority int) {
	item := &Item{
		value:    value,
		priority: priority,
	}
	pq.Items = append(pq.Items, item)
	sort.SliceStable(pq.Items, func(i, j int) bool {
		if pq.HeapTypeMin {
			return pq.Items[i].priority < pq.Items[j].priority
		}
		return pq.Items[i].priority > pq.Items[j].priority
	})
}

func (pq *PriorityQueue) Dequeue() (string, error) {
	if len(pq.Items) == 0 {
		return "", errors.New("PriorityQueue is empty")
	}
	item := (pq.Items)[0]
	pq.Items = (pq.Items)[1:]
	return item.value, nil
}

func (pq *PriorityQueue) Peek() (string, error) {
	if len(pq.Items) == 0 {
		return "", errors.New("PriorityQueue is empty")
	}
	item := (pq.Items)[0]
	return item.value, nil
}

func main() {
	pq := &PriorityQueue{HeapTypeMin: true}
	pq.Enqueue("Task 1", 2)
	pq.Enqueue("Task 2", 1)
	pq.Enqueue("Task 3", 3)
	peek, _ := pq.Peek()
	fmt.Println("Peek at top priority item:", peek)
	deq, _ := pq.Dequeue()
	fmt.Println("Dequeued:", deq)
}
