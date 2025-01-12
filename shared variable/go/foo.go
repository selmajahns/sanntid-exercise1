package main

import (
	"fmt"
	"runtime"
)

type action int

const (
	increment action = iota
	decrement
	get
)

type message struct {
	action action
	reply  chan int
}

func numberServer(ch chan message, done chan struct{}) {
	i := 0
	for {
		select {
		case msg := <-ch:
			switch msg.action {
			case increment:
				i++
			case decrement:
				i--
			case get:
				msg.reply <- i
			}
		case <-done:
			close(ch)
			return
		}
	}
}

func incrementing(ch chan message, done chan struct{}) {
	for j := 0; j < 1000000; j++ {
		ch <- message{action: increment}
	}
	done <- struct{}{}
}

func decrementing(ch chan message, done chan struct{}) {
	for k := 0; k < 1000001; k++ {
		ch <- message{action: decrement}
	}
	done <- struct{}{}
}

func main() {
	runtime.GOMAXPROCS(2)

	ch := make(chan message)
	done := make(chan struct{}, 2)
	go numberServer(ch, done)

	go incrementing(ch, done)
	go decrementing(ch, done)

	<-done
	<-done

	replyChan := make(chan int)
	ch <- message{action: get, reply: replyChan}
	finalValue := <-replyChan

	fmt.Println("The magic number is:", finalValue)
}
