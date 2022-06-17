package main

import (
	"fmt"

	"github.com/ZiXyos/goutils/pkg/linked_list"
)

func main() {

	list := linked_list.List{

		Head: nil,
		Tail: nil,
	}

	list.AddElemAtFront(10)
	list.AddElemAtFront(100)
	list.AddElemAtFront(1000)
	list.AddElemAtFront(10000)

	list.Dump()
	fmt.Print(list.GetSize())
}
