package main

import (
	"fmt"
	"sync"
)

func main() {
	{
		var mu sync.Mutex
		var wg sync.WaitGroup

		m := make(map[int]int, 10)

		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				for j := 10; j > 0; j-- {
					mu.Lock()
					m[i] = j
					mu.Unlock()
				}
			}(i)
		}

		wg.Wait()

		fmt.Println("Обычная map:", m)
	}
	{
		var sm sync.Map
		var wg sync.WaitGroup

		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				for j := 10; j > 0; j-- {
					sm.Store(i, j)
				}
			}(i)
		}

		wg.Wait()

		fmt.Print("sync map: ")
		sm.Range(func(k, v interface{}) bool {
			fmt.Printf("[%v:%v]", k, v)
			return true
		})
	}
}
