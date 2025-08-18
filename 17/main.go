package main

import (
	"fmt"
	"slices"
)

func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] > target {
			right = mid - 1
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}

	return -1
}

func main() {
	arr := []int{423, 342, 342, 3452, 53, 6754, 23, 342, 4321, 142, 6453, 32, 543, 213, 33, 547, 453}
	slices.Sort(arr)

	fmt.Println(binarySearch(arr, 33))
}
