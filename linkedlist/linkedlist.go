package linkedlist

import "fmt"
import "github.com/go_datastructure"

type Node struct {
	data interface{}
	next *Node
}

type LinkedList struct {
	head *Node
	size int
}

func New() *LinkedList {
	linkedlist := LinkedList{head: nil, size: 0}
	return &linkedlist
}

func NewNode(x interface{}) *Node {
	node := Node{data: x, next: nil}
	return &node
}

func (linklist *LinkedList) AddNode(x interface{}) interface{} {
	head := linklist.head
	newnode := NewNode(x)
	if head == nil {
		linklist.head = newnode
	} else {
		newnode.next = linklist.head
		linklist.head = newnode
	}
	linklist.size += 1
	return x
}

func (linklist *LinkedList) PrintList() {
	temp := linklist.head
	for temp != nil {
		fmt.Println(temp.data)
		temp = temp.next
	}
}

func (linklist *LinkedList) Size() int {
	return linklist.size
}

func (linklist *LinkedList) Empty() bool {
	return linklist.Size() == 0
}

func (linklist *LinkedList) SearchNode(x interface{}) (y interface{}, err error) {
	temp := linklist.head
	for temp != nil {
		if temp.data == y {
			return
		}
		temp = temp.next
	}
	err = go_datastructure.ErrorNotFound
	y = nil
	return
}

func (linklist *LinkedList) DeleteNode(x interface{}) (y interface{}, err error) {
	temp := linklist.head
	var prev *Node
	for temp != nil {
		if temp.data == x {
			//Node to be deleted is Head Node...
			if temp == linklist.head {
				linklist.head = linklist.head.next
				temp.next = nil
				temp = nil
			} else {
				if temp.next == nil {
					prev.next = nil
				} else {
					prev.next = temp.next
				}
				temp.next = nil
				temp = nil
			}
			linklist.size -= 1
			return
		}
		prev = temp
		temp = temp.next
	}
	y = nil
	err = go_datastructure.ErrorNotFound
	return
}
