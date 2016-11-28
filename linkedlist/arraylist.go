package linkedlist

import "math"
import "github.com/go_datastructure"

//import "fmt"
import "github.com/go_datastructure/common/log"

type ArrayList struct {
	buffer []interface{}
	size   int
}

const (
	LOAD_FACTOR   = 0.75
	INITIAL_CAP   = 10
	RESIZE_FACTOR = 1.5
)

func NewArrayList() *ArrayList {
	buf := make([]interface{}, INITIAL_CAP)
	list := ArrayList{buffer: buf}
	return &list
}

func (list *ArrayList) AddNode(x interface{}) interface{} {
	var actual_length int = list.length()
	log.InfoWithFields("Length Of List Before adding", log.Fields{"Length: ": actual_length})
	if actual_length >= int(math.Floor(float64(list.capacity())*LOAD_FACTOR)) {
		//Resize needs to Happen with 1.5 times
		log.InfoWithFields("Resize Operation with ", log.Fields{"Length: ": actual_length})
		dest := make([]interface{}, int(math.Floor(float64(list.capacity())*RESIZE_FACTOR)))
		copy(dest, list.buffer)
		list.buffer = nil
		list.buffer = dest
		log.InfoWithFields("Capacity of ArrayList", log.Fields{"Length: ": list.capacity()})
	}
	list.buffer[actual_length] = x
	list.size += 1
	return x
}

func (list *ArrayList) SearchNode(x interface{}) (y int, err error) {
	//Return the Index
	var len int = list.length()
	for i := 0; i < len; i++ {
		if list.buffer[i] == x {
			y = i
			return
		}
	}
	err = go_datastructure.ErrorNotFound
	y = -1
	return
}

func (list *ArrayList) DeleteNode(x interface{}) (y interface{}, err error) {
	//Delete first occurance of x
	var len int = list.length()
	for i := 0; i < len; i++ {
		if list.buffer[i] == x {
			if i+1 < len {
				//Concatenating Two Slices... Look for https://www.dotnetperls.com/slice-go
				list.buffer = append(list.buffer[:i], list.buffer[i+1:]...)
			} else {
				list.buffer = list.buffer[:i]
			}
			y = x
			list.size -= 1
			return
		}
	}
	err = go_datastructure.ErrorNotFound
	return
}

func (list *ArrayList) PrintList() {
	log.InfoWithFields("Length of ArrayList", log.Fields{"Length: ": list.length()})
	log.InfoWithFields("ArrayList", log.Fields{"List State ": list.buffer})
}

func (list *ArrayList) Empty() bool {
	return list.length() == 0
}

func (list *ArrayList) Size() int {
	return list.length()
}

func (list *ArrayList) length() int {
	return list.size
}

func (list *ArrayList) capacity() int {
	return len(list.buffer)
}
