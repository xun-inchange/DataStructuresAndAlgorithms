package array_test

import (
	"DataStructuresAndAlgorithms/Array"
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	var list Array.List = Array.NewArrayList()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	for i := 0; i < 20; i++ {
		list.Insert(1, i)
	}
	fmt.Println(list)
}
