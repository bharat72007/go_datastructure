// Copyright (c) 2016, Bharat,Jain. All rights reserved.
// Package stack implements a stack backed by Integer Slice.
// Structure is not thread safe.
// Reference: https://en.wikipedia.org/wiki/Stack_%28abstract_data_type%29#Array

package stack

import "github.com/go_datastructure"
import "github.com/go_datastructure/common/log"

//Structure Stack conatins Backed Slice of Integer, Reference Top of Stack and Current Size of Stack
type Stack struct {
	curr_size, top int
	elements       []interface{}
}

//Operations Contract for Stack
type Operations interface {
	Pop() (interface{}, error)
	Push(x interface{}) interface{}
	Peek() (interface{}, error)
}

//Pop Operation, where element is poped off from the top of Stack
func (stack *Stack) Pop() (top interface{}, err error) {
	log.InfoWithFields("Pop from Stack", log.Fields{"top": top})
	if stack.curr_size > 0 {
		top = stack.elements[stack.top]
		stack.elements = stack.elements[:stack.curr_size-1]
		stack.top -= 1
		stack.curr_size -= 1
		return
	}
	err = go_datastructure.ErrorUnderFlow
	return
}

//Push Operation, where element is added at the top of Stack
func (stack *Stack) Push(x interface{}) interface{} {
	log.InfoWithFields("Push into Stack", log.Fields{"x": x})
	stack.elements = append(stack.elements, x)
	stack.top += 1
	stack.curr_size += 1
	topEle := stack.elements[stack.top]
	return topEle
}

//Peek Operation, where Top element of Stack can be seen.
func (stack *Stack) Peek() (top interface{}, err error) {
	log.InfoWithFields("Peeking into Stack", log.Fields{"top": top})
	if stack.curr_size > 0 {
		top = stack.elements[stack.top]
		return
	}
	err = go_datastructure.ErrorUnderFlow
	return
}

//Initialize a Stack with initial values and return it
func New() Stack {
	//Initialization of Stack based using make
	log.InfoWithFields("Initializing Stack", log.Fields{})
	var stack Stack = Stack{curr_size: 0, top: -1, elements: make([]interface{}, 0, 0)}
	return stack
}
