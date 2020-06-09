package array

import (
	"DataStructuresAndAlgorithms/Array"
	"fmt"
	"testing"
)

func TestArrayIteratorStack(t *testing.T) {
	stackIt := Array.NewArrayListStackX()
	stackIt.PushX(1)
	stackIt.PushX(2)
	stackIt.PushX(3)
	stackIt.PushX(4)
	stackIt.PushX(5)
	stackIt.PushX(6)
	//stackIt.PopX()
	//stackIt.PopX()
	//stackIt.PopX()
	//stackIt.PopX()
	//stackIt.PopX()
	//stackIt.PopX()

	for it := stackIt.MyIt; it.HasNext(); {
		data, _ := it.Next()
		fmt.Println(data)
	}
}
