package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	nums := make(chan int)
	nums2 := make(chan int)

	wg.Add(2)
	go func() {
		defer close(nums)
		arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		for _, v := range arr {
			nums <- v
		}
	}()

	go func() {
		defer close(nums2)
		for v := range nums {
			nums2 <- v * 2
		}
	}()

	for v := range nums2 {
		fmt.Println(v)
	}
}
