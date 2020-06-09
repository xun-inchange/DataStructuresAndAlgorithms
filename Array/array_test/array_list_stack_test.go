package array

import (
	"DataStructuresAndAlgorithms/Array"
	"fmt"
	"testing"
)

func TestArrayListStack(t *testing.T) {
	stack := Array.NewArrayListStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	stack.Push(6)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())

}
