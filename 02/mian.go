package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	arr := []int{2, 4, 6, 7, 10}

	for _, v := range arr {
		wg.Add(1)

		go func(val int) {
			defer wg.Done()
			fmt.Println(val * val)
		}(v)
	}

	wg.Wait()
}
