package goaster

import (
	"testing"
)

func TestQueue_EnqueueAndDequeue(t *testing.T) {
	q := NewQueue()

	toast1 := NewToast("Toast 1", InfoLevel)
	toast2 := NewToast("Toast 2", SuccessLevel)

	q.Enqueue(toast1)
	q.Enqueue(toast2)

	if q.Size() != 2 {
		t.Errorf("expected size 2, got %d", q.Size())
	}

	item, err := q.Dequeue()
	if err != nil {
		t.Fatalf("unexpected error during dequeue: %v", err)
	}
	if item != toast1 {
		t.Errorf("expected first dequeued item to be toast1, got %+v", item)
	}

	item, err = q.Dequeue()
	if err != nil {
		t.Fatalf("unexpected error during dequeue: %v", err)
	}
	if item != toast2 {
		t.Errorf("expected second dequeued item to be toast2, got %+v", item)
	}

	_, err = q.Dequeue()
	if err == nil {
		t.Errorf("expected error on dequeue from empty queue, got nil")
	}
}

func TestQueue_IsEmpty(t *testing.T) {
	q := NewQueue()
	if !q.IsEmpty() {
		t.Error("expected queue to be empty")
	}

	q.Enqueue(NewToast("Hello", DefaultLevel))
	if q.IsEmpty() {
		t.Error("expected queue to be non-empty after enqueue")
	}
}

func TestQueue_Size(t *testing.T) {
	q := NewQueue()
	if q.Size() != 0 {
		t.Errorf("expected size 0, got %d", q.Size())
	}
	q.Enqueue(NewToast("One", InfoLevel))
	q.Enqueue(NewToast("Two", SuccessLevel))
	if q.Size() != 2 {
		t.Errorf("expected size 2, got %d", q.Size())
	}
}

func TestQueue_GetMessagesAndDequeue(t *testing.T) {
	q := NewQueue()
	toast1 := NewToast("Message A", WarningLevel)
	toast2 := NewToast("Message B", ErrorLevel)

	q.Enqueue(toast1)
	q.Enqueue(toast2)

	messages := q.GetMessagesAndDequeue()
	if len(messages) != 2 {
		t.Fatalf("expected 2 messages, got %d", len(messages))
	}
	if !q.IsEmpty() {
		t.Error("expected queue to be empty after GetMessagesAndDequeue")
	}
	if messages[0].Message != "Message A" || messages[1].Message != "Message B" {
		t.Error("messages are not in correct order")
	}
}

func TestQueue_String(t *testing.T) {
	q := NewQueue()
	q.Enqueue(NewToast("First", DefaultLevel))
	q.Enqueue(NewToast("Second", SuccessLevel))

	str := q.String()
	if str == "" || str[:6] != "Queue:" {
		t.Errorf("unexpected String output: %s", str)
	}
}
