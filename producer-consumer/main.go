package main

import (
	"fmt"
	"sync"
	"time"
)

var timeout = make(chan int, 1)
var start = time.Now()
var buffer = make([]string, 0)
var wg = sync.WaitGroup{}

func flushToDisk(buf []string) {
	fmt.Println("flushToDisk", buf)
}

func consumer(msg <-chan string) {
	for {
		select {
		case <-timeout:
			fmt.Println("timeout, buffer: ", buffer)
			flushToDisk(buffer)
			reset()
		default:
			fmt.Println("default, buffer: ", buffer)
			if len(buffer) < 100 {
				buffer = append(buffer, <-msg)
			} else {
				flushToDisk(buffer)
				reset()
			}
		}
		time.Sleep(1 * time.Second)
	}
}

func producer(msg chan<- string) {
	defer wg.Done()
	for {
		time.Sleep(1 * time.Second)
		msg <- "aaa"
	}
}

func reset() {
	fmt.Println("reset")
	buffer = make([]string, 0)
	start = time.Now()
}

func timer() {
	for {
		if time.Now().Sub(start).Seconds() > 5 {
			timeout <- 1
		}
	}
}

func main() {
	msg := make(chan string)

	wg.Add(1)

	go consumer(msg)
	go producer(msg)
	go timer()

	wg.Wait()
}
