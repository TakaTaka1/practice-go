package main

import (
	"fmt"
)

type findValue struct {
	arr  []int
	find int
}

func binary(arr []int, find int) bool {
	left := 0
	right := int(len(arr)) - 1
	for left <= right {
		mid := (left + right) / 2
		if find < arr[mid] {
			right = mid - 1
		}
		if find > arr[mid] {
			left = mid + 1
		}
		if find == arr[mid] {
			return true
		}
	}
	return false
}

func main() {
	var forBinary findValue
	forBinary.arr = []int{2, 3, 4, 5, 6, 8}
	forBinary.find = 9
	fmt.Println(binary(forBinary.arr, forBinary.find)) // false
}
