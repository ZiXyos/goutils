package linkedlist

import (
	"errors"
	"fmt"
)

type LinkedList struct {
	data interface{}
	next *LinkedList
}

func Init(data interface{}) *LinkedList {
	return &LinkedList{

		data: data,
		next: nil,
	}
}

func (list *LinkedList) GetSize() int {

	var i int = 0

	if list.IsEmpty() {

		return i
	}

	//tmp := (*list)
	fmt.Println(list)

	return i
}

func (list *LinkedList) IsEmpty() bool {

	if list != nil {

		return false
	} else {

		return true
	}
}

func (list *LinkedList) Print() (bool, error) {

	if list.IsEmpty() {

		fmt.Println("list is empty")
		return false, nil
	}

	//tmp := list

	fmt.Println(list)

	return true, nil
}

func (list *LinkedList) AddElemAtFront(_data interface{}) {

	if _data == nil {

		return
	}

	tmp := LinkedList{

		data: _data,
		next: list,
	}

	(*list) = tmp
}

func (list *LinkedList) AddElemtAtBack(data interface{}) (bool, error) {

	if data == nil {

		return false, errors.New("no data")
	}

	var tmp *LinkedList = &LinkedList{

		data: data,
		next: nil,
	}

	if list.IsEmpty() {

		list = tmp
		return true, nil
	}

	var last *LinkedList = list

	for !last.next.IsEmpty() {

		//log.Println(last.data)
		last = last.next
	}

	last.next = tmp

	return true, nil
}
