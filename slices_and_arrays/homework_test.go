package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type CircularQueue struct {
	values []int
	// need to implement
	size, rear, front int // two pointers
}

func NewCircularQueue(size int) CircularQueue {
	return CircularQueue{
		values: make([]int, size),
		// need to implement
		size:  size,
		rear:  -1,
		front: -1,
	}
}

func (q *CircularQueue) Push(value int) bool {
	// need to implement
	if q.Full() {
		return false
	}

	if q.Empty() {
		q.front = 0
	}

	q.rear = (q.rear + 1) % q.size
	q.values[q.rear] = value

	return true
}

func (q *CircularQueue) Pop() bool {
	// need to implement
	if q.Empty() {
		return false
	}

	// for the last element
	if q.rear == q.front {
		q.rear = -1
		q.front = -1
		return true
	}

	q.front = (q.front + 1) % q.size

	return true
}

func (q *CircularQueue) Front() int {
	if q.front == -1 {
		return -1
	}
	return q.values[q.front] // need to implement
}

func (q *CircularQueue) Back() int {
	if q.rear == -1 {
		return -1
	}
	return q.values[q.rear] // need to implement
}

func (q *CircularQueue) Empty() bool {
	// need to implement
	return q.rear == -1 && q.front == -1
}

func (q *CircularQueue) Full() bool {
	// need to implement
	return q.front == q.size-1 || q.front == (q.rear+1)%q.size
}

func TestFull(t *testing.T) {
	queue := NewCircularQueue(3)
	queue.values = []int{1, 2, 3}
	queue.front = 0
	queue.rear = 2

	assert.True(t, queue.Full())

	queue.front = 1
	queue.rear = 0

	assert.True(t, queue.Full())
}

func TestCircularQueue(t *testing.T) {
	const queueSize = 3
	queue := NewCircularQueue(queueSize)

	assert.True(t, queue.Empty())
	assert.False(t, queue.Full())

	assert.Equal(t, -1, queue.Front())
	assert.Equal(t, -1, queue.Back())
	assert.False(t, queue.Pop())

	assert.True(t, queue.Push(1))
	assert.True(t, queue.Push(2))
	assert.True(t, queue.Push(3))
	assert.False(t, queue.Push(4))

	assert.True(t, reflect.DeepEqual([]int{1, 2, 3}, queue.values))

	assert.False(t, queue.Empty())
	assert.True(t, queue.Full())

	assert.Equal(t, 1, queue.Front())
	assert.Equal(t, 3, queue.Back())

	assert.True(t, queue.Pop())
	assert.False(t, queue.Empty())
	assert.False(t, queue.Full())
	assert.True(t, queue.Push(4))

	assert.True(t, reflect.DeepEqual([]int{4, 2, 3}, queue.values))

	assert.Equal(t, 2, queue.Front())
	assert.Equal(t, 4, queue.Back())

	assert.True(t, queue.Pop())
	assert.True(t, queue.Pop())
	assert.True(t, queue.Pop())
	assert.False(t, queue.Pop())

	assert.True(t, queue.Empty())
	assert.False(t, queue.Full())
}

func TestCircularQueueAdditional(t *testing.T) {
	const queueSize = 3
	queue := NewCircularQueue(queueSize)
	queue.values = []int{1, 0, 3}
	queue.front = 2
	queue.rear = 0
	assert.True(t, queue.Pop())
	assert.Equal(t, 0, queue.front)
}
