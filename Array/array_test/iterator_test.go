package array

import (
	"DataStructuresAndAlgorithms/Array"
	"fmt"
	"testing"
)

func TestIterator(t *testing.T) {
	var list Array.List = Array.NewArrayList()
	list.Append("a1")
	list.Append("b2")
	list.Append("c3")
	for it := list.NewListIterator(); it.HasNext(); {
		item, _ := it.Next()
		if item == "b2" {
			it.Remove()
		}
		fmt.Println(item)
	}
	fmt.Println(list)
}
