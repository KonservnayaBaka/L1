package main

import "fmt"

func partition(arr []int, left, right int) int {
	mid := (left + right) / 2
	pivot := arr[mid]

	arr[mid], arr[right] = arr[right], arr[mid]

	i := left - 1
	for j := left; j < right; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[right] = arr[right], arr[i+1]
	return i + 1
}

func quickSort(arr []int, left, right int) {
	if left < right {
		pi := partition(arr, left, right)
		//left
		quickSort(arr, left, pi-1)
		//right
		quickSort(arr, pi+1, right)
	}
}

func QuickSort(arr []int) {
	if len(arr) > 1 {
		quickSort(arr, 0, len(arr)-1)
	}
}

func main() {
	arr := []int{10, 7, 8, 9, 1, 5}
	QuickSort(arr)

	fmt.Println("Отсортированный массив:", arr)
}
