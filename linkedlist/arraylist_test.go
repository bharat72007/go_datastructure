package linkedlist

import "testing"

func TestAddToList(t *testing.T) {
	arrlist := NewArrayList()
	arrlist.AddNode(10)
	arrlist.AddNode(30)
	arrlist.AddNode(40)
	arrlist.AddNode(50)
	arrlist.AddNode(60)
	arrlist.AddNode(70)
	arrlist.AddNode(80)
	arrlist.AddNode(90)
	arrlist.PrintList()

	arrlist.DeleteNode(80)
	arrlist.PrintList()
	arrlist.AddNode(80)
	arrlist.AddNode(100)
	arrlist.AddNode(110)
	arrlist.PrintList()
	arrlist.DeleteNode(110)
	arrlist.PrintList()
	arrlist.DeleteNode(10)
	arrlist.PrintList()
}
