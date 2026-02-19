package goaster

import (
	"errors"
	"fmt"
	"strings"
)

// ErrEmptyQueue is returned when attempting to dequeue from an empty queue.
var ErrEmptyQueue = errors.New("queue is empty")

// Queue is a FIFO container for Toast notifications.
type Queue struct {
	elements []Toast
}

// NewQueue creates and returns an empty Queue.
func NewQueue() *Queue {
	return &Queue{elements: make([]Toast, 0)}
}

// Enqueue adds an element to the end of the queue.
func (q *Queue) Enqueue(toast Toast) {
	q.elements = append(q.elements, toast)
}

// Dequeue removes and returns the first element of the queue.
// If the queue is empty, it returns a zero-value Toast and ErrEmptyQueue.
func (q *Queue) Dequeue() (Toast, error) {
	if len(q.elements) == 0 {
		return Toast{}, ErrEmptyQueue
	}
	element := q.elements[0]
	q.elements[0] = Toast{} // zero the slot to allow GC
	q.elements = q.elements[1:]
	return element, nil
}

// IsEmpty checks if the queue is empty.
func (q *Queue) IsEmpty() bool {
	return len(q.elements) == 0
}

// Size returns the number of elements in the queue.
func (q *Queue) Size() int {
	return len(q.elements)
}

// String returns a string representation of the queue.
func (q *Queue) String() string {
	var sb strings.Builder
	sb.WriteString("Queue: front [")
	for i, e := range q.elements {
		fmt.Fprintf(&sb, "%v", e)
		if i < len(q.elements)-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString("] rear")
	return sb.String()
}

// DrainAll removes and returns all Toast messages from the queue.
// After this call the queue is empty.
func (q *Queue) DrainAll() []Toast {
	messages := q.elements
	q.elements = nil
	return messages
}
