package main

import "fmt"

/*
给定一个已按照 升序排列  的整数数组 numbers ，请你从数组中找出两个数满足相加之和等于目标数 target 。
函数应该以长度为 2 的整数数组的形式返回这两个数的下标值。numbers 的下标 从 1 开始计数

范例：
输入：numbers = [2,7,11,15], target = 9
输出：[1,2]
解释：2 与 7 之和等于目标数 9 。因此 index1 = 1, index2 = 2 。
*/

func Test(numbers []int, target int) interface{} {
	low := 0
	high := len(numbers) - 1
	for low <= high {
		if target == numbers[low]+numbers[high] {
			return []int{low + 1, high + 1}
		} else if target < numbers[low]+numbers[high] {
			high--
		} else {
			low++
		}
	}
	return []int{-1, -1}
}

func main() {
	numbers := []int{2, 7, 11, 15}
	res := Test(numbers, 9)
	fmt.Println(res)
}
