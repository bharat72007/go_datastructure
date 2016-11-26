package linkedlist

type Operations interface {
	AddNode(interface{}) interface{}
	SearchNode(interface{}) (interface{}, error)
	DeleteNode(interface{}) (interface{}, error)
	Empty() bool
	Size() int
	PrintList()
}
