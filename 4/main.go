package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func workerPull(nums <-chan int, res chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range nums {
		res <- n * n
	}
}

func main() {
	var countWorkers int
	fmt.Print("Введите количество воркеров: ")
	fmt.Scan(&countWorkers)

	ctx, cancel := context.WithCancel(context.Background())
	nums := make(chan int, countWorkers)
	res := make(chan int, countWorkers)
	quit := make(chan os.Signal, 1)

	var wg sync.WaitGroup

	for i := 0; i < countWorkers; i++ {
		wg.Add(1)
		go workerPull(nums, res, &wg)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		num := 0
		for {
			select {
			case <-ctx.Done():
				return
			default:
				nums <- num
				num++
				fmt.Println(<-res)
				time.Sleep(time.Second)
			}
		}
	}()

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("Shooting down...")
	cancel()
	close(nums)
	wg.Wait()
}
