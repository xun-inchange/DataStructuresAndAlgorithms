package main

import "fmt"

func main() {
	arr := []int{1, 2, 4, 9, 7, 6, 5}
	insertSort(arr)
	fmt.Println(arr)
}

func insertSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			} else {
				break
			}
		}
	}
}
