package main

import "fmt"

/*快速排序*/

func main() {
	arr := []int{3, 9, 2, 8, 1, 7, 4, 6, 5, 10}
	newArr := quickSort(arr)
	fmt.Println(newArr)
}

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	//n := len(arr) - 1
	n := 0
	splitData := arr[n] //分割数据
	lowArr := make([]int, 0)
	highArr := make([]int, 0)
	midArr := make([]int, 0)
	midArr = append(midArr, splitData)
	for i := 0; i < len(arr); i++ {
		if i == n {
			continue
		}
		if arr[i] < splitData {
			lowArr = append(lowArr, arr[i])
		} else if arr[i] > splitData {
			highArr = append(highArr, arr[i])
		} else {
			midArr = append(midArr, arr[i])
		}
	}
	lowArr, highArr = quickSort(lowArr), quickSort(highArr)
	newArr := append(append(lowArr, midArr...), highArr...)
	return newArr
}
