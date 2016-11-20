package queue

import "github.com/go_datastructure"

type Queue struct {
	front, rear int
	buffer      []int
}

type Operations interface {
	Enque(int) int
	Deque() (int, error)
	Peek() (int, error)
	Empty() bool
	QueueSize() int
}

func New() Queue {
	queue := Queue{front: -1, rear: -1, buffer: make([]int, 0, 0)}
	return queue
}

func (queue *Queue) Empty() bool {
	if queue.QueueSize() == 0 {
		return true
	}
	return false
}

func (queue *Queue) QueueSize() int {
	return queue.rear - queue.front
}

func (queue *Queue) Peek() (int, error) {
	if queue.Empty() {
		return 0, go_datastructure.ErrorUnderFlow
	}
	return queue.buffer[0], nil
}

func (queue *Queue) Enque(x int) int {
	queue.buffer = append(queue.buffer, x)
	if queue.Empty() {
		queue.front = 0
		queue.rear = 1
	} else {
		queue.rear += 1
	}
	return queue.buffer[0]
}

func (queue *Queue) Deque() (int, error) {
	if queue.Empty() {
		return 0, go_datastructure.ErrorUnderFlow
	}
	if queue.QueueSize() == 1 {
		queue.front = -1
		queue.rear = -1
	} else {
		queue.rear -= 1
	}
	frontEle := queue.buffer[0]
	queue.buffer = queue.buffer[1:]
	return frontEle, nil
}
