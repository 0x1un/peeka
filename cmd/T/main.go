package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		// time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	N := 10
	exit := make(chan struct{})
	done := make(chan struct{}, N)
	// start N worker goroutines
	for i := 0; i < N; i++ {
		go func(n int) {
			for {
				select {
				// wait for exit signal
				case <-exit:
					fmt.Printf("worker goroutine #%d exit\n", n)
					done <- struct{}{}
					return
				case <-time.After(time.Second):
					fmt.Printf("worker goroutine #%d is working...\n", n)
				}
			}
		}(i)
	}
	time.Sleep(3 * time.Second)
	// broadcast exit signal
	close(exit)
	// wait for all worker goroutines exit
	for i := 0; i < N; i++ {
		<-done
	}
	fmt.Println("main goroutine exit")
}
