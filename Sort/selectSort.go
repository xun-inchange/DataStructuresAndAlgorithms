package main

import "fmt"

func main() {
	arr := []int{1, 9, 2, 8, 7, 6, 4, 5}
	selectSort(arr)
	fmt.Println(arr)
}

func selectSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ { //因为最后一个元素已经可以不用选择了，它就是排在最后的
		min := i // 标记索引
		for j := i + 1; j < len(arr); j++ {
			//如果arr[i] < arr[j]
			if arr[min] < arr[j] {
				min = j //保存极大值的索引
			}
		}
		//如果min不等于i 那么就交换
		if min != i {
			arr[i], arr[min] = arr[min], arr[i]
		}
	}
}
