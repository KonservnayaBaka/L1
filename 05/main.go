package main

import (
	"fmt"
	"sync"
	"time"
)

func printNumFromChanel(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range ch {
		fmt.Println(n)
	}
}

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup

	wg.Add(2)
	go printNumFromChanel(ch, &wg)

	go func() {
		defer wg.Done()
		num := 0
		timeout := time.After(5 * time.Second)
		for {
			select {
			case <-timeout:
				close(ch)
				return
			default:
				ch <- num
				num++
				time.Sleep(time.Millisecond * 500)
			}
		}
	}()

	wg.Wait()
	fmt.Println("Timeout")
}
