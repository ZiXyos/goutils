package linked_list

import (
	"testing"
)

func AddElemAtFront_T(t *testing.T) {

	list := List{}
	if &list == nil {

		t.Error("list not created")
	}

	list.AddElemAtFront("1st injected")
	list.AddElemAtFront("2nd injected")
}
