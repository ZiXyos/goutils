package main

import (
	"github.com/ZiXyos/goutils/pkg/linked_list"
)

func main() {

	list := linked_list.List{

		Head: nil,
		Tail: nil,
	}

	list.AddElemAtFront(10)
	list.AddElemtAtBack(100)
	list.AddElemtAtBack(1000)
	list.AddElemAtFront(0)
	list.AddElemAtFront("list")

	list.Debug()
	list.Dump()
}
