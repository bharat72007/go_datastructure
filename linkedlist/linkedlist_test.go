package linkedlist

import "testing"
import "fmt"

func TestAddOnly(t *testing.T) {
	linklist := New()
	linklist.Add("a")
	linklist.Add("b")
	linklist.Add("c")
	//	linklist.PrintList()
}

func TestDeleteNode(t *testing.T) {
	linklist := New()
	linklist.Add("a")
	linklist.Add("b")
	linklist.Add("c")
	linklist.DeleteNode("c")
	fmt.Println()
	linklist.PrintList()
	linklist.DeleteNode("a")
	fmt.Println()
	linklist.PrintList()
	linklist.DeleteNode("b")
	fmt.Println()
	linklist.PrintList()
	_, ok := linklist.DeleteNode("b")
	if ok != nil {
		t.Errorf("Linked List Element not present")
	}

}
