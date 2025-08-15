package main

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(7)
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	go func() {
		defer wg.Done()
		fmt.Println("Завершение по окончанию работы")
		return
	}()
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	ch := make(chan int)
	go func() {
		defer wg.Done()
		for v := range ch {
			_ = v
		}
		fmt.Println("Завершение по закрытию канала")
	}()
	ch <- 6
	close(ch)
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	ctxCancel, cancel := context.WithCancel(context.Background())

	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctxCancel.Done():
				fmt.Println("Завершение по контексту")
				return
			default:
				time.Sleep(time.Second)
			}
		}
	}()

	cancel()
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	go func() {
		defer wg.Done()
		for {
			select {
			case <-time.After(time.Second):
				fmt.Println("Завершение по таймауту (time after)")
			default:
				time.Sleep(time.Millisecond * 500)
			}
		}
	}()

	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second)

	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctxTimeout.Done():

				fmt.Println("Завершение по таймауту (context with timeout)")
				return
			default:
				time.Sleep(time.Millisecond * 500)
			}
		}
	}()

	cancel()
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	go func() {
		defer wg.Done()
		defer fmt.Println("Завершение по runtime.Goexit")
		runtime.Goexit()
	}()
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	go func() {
		defer wg.Done()
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Завершение по панике")
			}
		}()
		panic("Завершение по панике")
	}()
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	quit := make(chan struct{})

	go func() {
		defer wg.Done()
		for {
			select {
			case <-quit:
				fmt.Println("Завершение по quit каналу")
				return
			default:
				time.Sleep(time.Millisecond * 500)
			}
		}
	}()
	quit <- struct{}{}
	close(quit)
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	wg.Wait()
	fmt.Println("Завершение по os.Exit(0) (Завершает абсолютно всё)")
	os.Exit(0)
}
