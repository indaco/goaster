package goaster

import (
	"fmt"
	"strings"
)

type Queue struct {
	elements []Toast
}

func NewQueue() *Queue {
	return &Queue{elements: make([]Toast, 0)}
}

// Enqueue adds an element to the end of the queue.
func (q *Queue) Enqueue(toast Toast) {
	q.elements = append(q.elements, toast)
}

// Dequeue removes and returns the first element of the queue.
// If the queue is empty, it returns -1 and an error.
func (q *Queue) Dequeue() (Toast, error) {
	if len(q.elements) == 0 {
		return Toast{}, fmt.Errorf("queue is empty")
	}
	element := q.elements[0]
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
		sb.WriteString(fmt.Sprintf("%v", e))
		if i < len(q.elements)-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString("] rear")
	return sb.String()
}

// GetMessagesAndDequeue returns a slice of Toast messages from the queue.
// It dequeues messages from the queue until it's empty.
// **NOTE**: in a-h/templ only `for...range` is a valid `for`loop.
func (q *Queue) GetMessagesAndDequeue() []Toast {
	var messages []Toast

	for q.Size() > 0 {
		item, err := q.Dequeue()
		if err != nil {
			fmt.Println("Error:", err)
			break
		}
		messages = append(messages, item)
	}
	return messages
}
