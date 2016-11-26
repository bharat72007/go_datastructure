// Copyright (c) 2016, Bharat,Jain. All rights reserved.
// Package queue implements a queue backed by Integer Slice.
// Structure is not thread safe.
package queue

import "github.com/go_datastructure"
import "github.com/go_datastructure/common/log"

//Defining strcut Queue which maintains, front and rear index and Holds a Integer Buffer
type Queue struct {
	front, rear int
	buffer      []interface{}
}

//Availble Contract Operations for Queue Structure
type Operations interface {
	Enque(interface{}) interface{}
	Deque() (interface{}, error)
	Peek() (interface{}, error)
	Empty() bool
	QueueSize() interface{}
}

//Initializing Queue with Default Values.
func New() Queue {
	//Making Use of Slice, Initialized similar to Array without mentioning size
	var buf []interface{}
	queue := Queue{front: -1, rear: -1, buffer: buf}
	return queue
}

//Check Whether Queue is Empty or Not
func (queue *Queue) Empty() bool {
	log.InfoWithFields("Check is Queue Empty", log.Fields{})
	if queue.QueueSize() == 0 {
		return true
	}
	return false
}

//Return number of elements in Queue
func (queue *Queue) QueueSize() int {
	return queue.rear - queue.front
}

//Peek the Front Element in Queue, If Queue is Empty return Error
func (queue *Queue) Peek() (front interface{}, err error) {
	log.InfoWithFields("Peek into Queue", log.Fields{"fron": front})
	if queue.Empty() {
		err = go_datastructure.ErrorUnderFlow
		return
	}
	front = queue.buffer[0]
	return
}

//Enque Element to the end of the Queue
func (queue *Queue) Enque(x interface{}) interface{} {
	log.InfoWithFields("Enque Operation", log.Fields{"x": x})
	queue.buffer = append(queue.buffer, x)
	if queue.Empty() {
		queue.front = 0
		queue.rear = 1
	} else {
		queue.rear += 1
	}
	return queue.buffer[0]
}

//Deque Element from Front of the Queue, If Queue is Empty return Error
func (queue *Queue) Deque() (front interface{}, err error) {
	log.InfoWithFields("Dequqe from Queue", log.Fields{"front": front})
	if queue.Empty() {
		err = go_datastructure.ErrorUnderFlow
		return
	}
	if queue.QueueSize() == 1 {
		queue.front = -1
		queue.rear = -1
	} else {
		queue.rear -= 1
	}
	front = queue.buffer[0]
	queue.buffer = queue.buffer[1:]
	return
}
