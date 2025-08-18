package main

import "fmt"

func main() {
	arr := []int{1, 2, 3}
	arr2 := []int{2, 3, 4}
	arr3 := []int{}
	m := map[int]bool{}

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr2); j++ {
			if arr[i] == arr2[j] {
				if !m[arr[i]] {
					arr3 = append(arr3, arr[i])
					m[arr[i]] = true
				}
			}
		}
	}

	fmt.Println(arr3)
}
