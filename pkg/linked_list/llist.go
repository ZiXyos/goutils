package linked_list

import (
	"errors"
	"fmt"
)

type Node struct {
	_data interface{}
	_next *Node
	_prev *Node
}

type List struct {
	Head *Node
	Tail *Node
}

func (list *Node) IsEmpty() bool {

	return list == nil
}

func (list *List) GetSize() int {

	var i int = 0
	if list == nil {

		return i
	}

	tmp := list
	var node *Node = tmp.Head

	for ; !node._next.IsEmpty(); i++ {

		node = node._next
	}

	return i + 1
}

func (list *List) Dump() error {

	if list == nil {

		return errors.New("list uninitialized")
	}

	var node *Node = list.Head
	for i := 0; !node.IsEmpty(); i++ {

		fmt.Printf("[%d]: %d \n", i, node._data)
		node = node._next
	}
	return nil
}

func (list *List) AddElemAtFront(data interface{}) error {

	if data == nil || list == nil {

		return errors.New("error at init")
	}

	tmp := Node{_data: data}

	if list.Head != nil {

		tmp._next = list.Head
		tmp._prev = nil
		list.Head = (&tmp)
	} else {

		list.Head = &tmp
	}

	return nil
}
