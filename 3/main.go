package main

import (
	"fmt"
	"time"
)

func workerPull(id int, ch chan int) {
	for n := range ch {
		fmt.Printf("worker %d: %d\n", id, n)
	}
}

func main() {
	var workersCount int
	fmt.Print("Введите количество воркеров: ")
	fmt.Scan(&workersCount)

	ch := make(chan int, workersCount)

	for i := 1; i <= workersCount; i++ {
		go workerPull(i, ch)
	}

	num := 0
	for {
		ch <- num
		num++
		time.Sleep(time.Second)
	}
}
