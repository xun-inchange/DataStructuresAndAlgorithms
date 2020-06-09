package StackArray

import "fmt"

//栈实现低级递归

func Add(num int) int {
	if num == 0 {
		return 0
	}
	return num + Add(num-1)
}

func stackRecursive() int {
	stack := NewStack()
	stack.Push(5)
	result := 0
	for !stack.IsEmpty() {
		data := stack.Pop() //取数据
		fmt.Println(stack.IsEmpty())
		if data == 0 {
			result += 0
		} else {
			result += data.(int)
			stack.Push(data.(int) - 1)
		}

	}
	return result
}
