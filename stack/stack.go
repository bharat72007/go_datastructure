// Copyright (c) 2016, Bharat,Jain. All rights reserved.
// Package stack implements a stack backed by Integer Slice.
// Structure is not thread safe.
// Reference: https://en.wikipedia.org/wiki/Stack_%28abstract_data_type%29#Array

package stack

import "github.com/go_datastructure"

//Structure Stack conatins Backed Slice of Integer, Reference Top of Stack and Current Size of Stack
type Stack struct {
	curr_size, top int
	elements       []int
}

//Operations Contract for Stack
type Operations interface {
	Pop() int
	Push(x int) int
	Peek() int
}

//Pop Operation, where element is poped off from the top of Stack
func (stack *Stack) Pop() (int, error) {
	if stack.curr_size > 0 {
		topEle := stack.elements[stack.top]
		stack.elements = stack.elements[:stack.curr_size-1]
		stack.top -= 1
		stack.curr_size -= 1
		return topEle, nil
	}
	return 0, go_datastructure.ErrorUnderFlow
}

//Push Operation, where element is added at the top of Stack
func (stack *Stack) Push(x int) int {
	stack.elements = append(stack.elements, x)
	stack.top += 1
	stack.curr_size += 1
	topEle := stack.elements[stack.top]
	return topEle
}

//Peek Operation, where Top element of Stack can be seen.
func (stack *Stack) Peek() (int, error) {
	if stack.curr_size > 0 {
		topEle := stack.elements[stack.top]
		return topEle, nil
	}
	return 0, go_datastructure.ErrorUnderFlow
}

//Initialize a Stack with initial values and return it
func New() Stack {
	var stack Stack = Stack{curr_size: 0, top: -1, elements: make([]int, 0, 0)}
	return stack
}
