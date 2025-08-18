package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    *sync.Mutex
	value int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
}

func (c *Counter) PrintValue() {
	fmt.Println(c.value)
}

func main() {
	var mu sync.Mutex
	var wg sync.WaitGroup

	counter := &Counter{
		mu:    &mu,
		value: 0,
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait()
	counter.PrintValue()
}
