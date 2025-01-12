package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		ch <- i
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("[producer]: pushing %d\n", i)
	}
	close(ch)
}

func consumer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	for i := range ch {
		fmt.Printf("[consumer]: %d\n", i)
		time.Sleep(50 * time.Millisecond)
	}

}

func main() {
	buffer := make(chan int, 10)
	var wg sync.WaitGroup
	wg.Add(2)
	// TODO: make a bounded buffer

	go consumer(buffer, &wg)
	go producer(buffer, &wg)

	//select {}
	wg.Wait()
}
