package linkedlist

import "github.com/go_datastructure"
import "fmt"

type DNode struct {
	data interface{}
	prev *DNode
	next *DNode
}

type DLinkedList struct {
	head *DNode
	size int
}

func (dlinklist *DLinkedList) AddNode(x interface{}) interface{} {
	new_node := New_Dnode(x)
	//temp_node := dlinklist.head
	if dlinklist.head == nil {
		dlinklist.head = new_node
	} else {
		new_node.next = dlinklist.head
		dlinklist.head.prev = new_node
		dlinklist.head = new_node //Beacuse the Type of DlinkedList.head is (*DNode) and of new_node is (*DNode) therefore is valid assignment.
		//If new_node doesn't return Pointer, then we might need to do ... dlinklist.head = &new_node
	}
	dlinklist.size += 1
	return x
}

func (dlinkedlist *DLinkedList) DeleteNode(x interface{}) (y interface{}, err error) {
	temp := dlinkedlist.head
	for temp != nil {
		if temp.data == x {
			y = x
			if temp == dlinkedlist.head {
				dlinkedlist.head = temp.next
				temp = nil
			} else {
				if temp.next == nil {
					temp.prev.next = nil
				} else {
					temp.prev.next = temp.next
					temp.next.prev = temp.prev
				}
				temp = nil
			}
			dlinkedlist.size -= 1
			return
		}
		temp = temp.next
	}
	err = go_datastructure.ErrorNotFound
	return
}

func (dlinklist *DLinkedList) Size() int {
	return dlinklist.size
}

func (dlinklist *DLinkedList) Empty() bool {
	return dlinklist.Size() == 0
}

func (dlinklist *DLinkedList) SearchNode(x interface{}) (y interface{}, err error) {
	temp := dlinklist.head
	for temp != nil {
		if x == temp.data {
			y = x
			return
		}
		temp = temp.next
	}
	err = go_datastructure.ErrorNotFound
	return
}

func (dlinklist *DLinkedList) PrintList() {
	temp := dlinklist.head
	for temp != nil {
		fmt.Println(temp.data)
		temp = temp.next
	}
}

func New_Dnode(x interface{}) *DNode {
	node := DNode{data: x, prev: nil, next: nil}
	return &node
}

func New_DLL() *DLinkedList {
	dlinklist := DLinkedList{head: nil, size: 0}
	return &dlinklist
}
