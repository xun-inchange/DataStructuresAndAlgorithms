package main

import "fmt"

func main() {
	arr := []int{1, 2, 4, 9, 7, 6, 4, 5}
	bubbleSort(arr)
	fmt.Println(arr)
}

func bubbleSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	for i := 0; i < len(arr)-1; i++ { //只剩一个就不要冒泡了
		isNeedExchange := false
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				isNeedExchange = true
			}
		}
		//两两比较的情况下，如果没有发生交换，说明不需要再进行冒泡排序了
		if !isNeedExchange {
			break
		}
	}
}
