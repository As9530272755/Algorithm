package main

import "fmt"

func BinarySearch(arr *[6]int, left int, right int, find int) {

	if left > right {
		fmt.Println("未找到", find)
		return
	}

	mid := (left + right) / 2

	if find < (*arr)[mid] {
		BinarySearch(arr, left, mid-1, find)
	}
	if find > (*arr)[mid] {
		BinarySearch(arr, mid+1, right, find)
	}
	if find == (*arr)[mid] {
		fmt.Println("找到", (*arr)[mid])
	}
}

func main() {
	arr := [6]int{1, 20, 22, 40, 51, 60}
	BinarySearch(&arr, 0, len(arr)-1, 60)
}
