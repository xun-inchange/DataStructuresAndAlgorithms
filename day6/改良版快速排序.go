package main

import (
	"fmt"
	"math/rand"
)

//快速排序思想：随机一个数字，比我小的往左，比我大的往右

func main() {
	arr := []int{3, 9, 2, 8, 1, 7, 4, 6, 5, 10}
	recursionQuickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

//插入排序
func insertSort(arr []int, left int, right int) {
	for i := left; i <= right; i++ { //迭代
		temp := arr[i] //每次取出临时变量
		var j int
		for j = i; j > left && arr[j-1] > temp; j-- { //定位
			arr[j] = arr[j-1] //数据往后移
		}
		arr[j] = temp //插入数据
	}
}

//递归快速排序
func recursionQuickSort(arr []int, left int, right int) {
	if right-left < 3 { //列表长度小于3的时候
		insertSort(arr, left, right) //进行插入排序
	} else {
		swap(arr, left, rand.Int()%(right-left+1)+left) //随机一个数，放在第一个位置
		number := arr[left]                             //坐标数字，比我小往后边，比我大往左边
		lt := left                                      //arr[left+1,lt] < number
		gt := right + 1                                 //arr[gt,right] > number
		i := left + 1                                   //arr[lt+1,i] == number
		for i < gt {
			if arr[i] < number {
				swap(arr, i, lt+1) //移动到小于的地方
				lt++
				i++
			} else if arr[i] > number {
				swap(arr, i, gt-1) //移动到大于的地方
				gt--

			} else {
				i++
			}
		}
		swap(arr, left, lt)                 //交换头部位置
		recursionQuickSort(arr, left, lt-1) //递归处理小于那一段
		recursionQuickSort(arr, gt, right)  //递归处理大于那一段

	}

}

//数据交换
func swap(arr []int, i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
