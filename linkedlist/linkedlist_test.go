package linkedlist

import "testing"
import "fmt"

func TestAddOnly(t *testing.T) {
	linklist := New()
	linklist.AddNode("a")
	linklist.AddNode("b")
	linklist.AddNode("c")
	//	linklist.PrintList()
}

func TestDeleteNode(t *testing.T) {
	linklist := New()
	linklist.AddNode("a")
	linklist.AddNode(1)
	linklist.AddNode("c")
	linklist.DeleteNode("c")
	fmt.Println()
	linklist.PrintList()
	linklist.DeleteNode("a")
	fmt.Println()
	linklist.PrintList()
	linklist.DeleteNode(1)
	fmt.Println()
	linklist.PrintList()
	_, ok := linklist.DeleteNode("b")
	if ok != nil {
		t.Errorf("Linked List Element not present")
	}

}
