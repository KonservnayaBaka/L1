package main

import (
	"fmt"
	"time"
)

func sleep(duration time.Duration) {
	timer := time.NewTimer(duration)
	<-timer.C
}

func sleepAfter(duration time.Duration) {
	<-time.After(duration)
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		sleep(1 * time.Second)
	}
	for i := 10; i < 20; i++ {
		fmt.Println(i)
		sleepAfter(1 * time.Second)
	}
}
