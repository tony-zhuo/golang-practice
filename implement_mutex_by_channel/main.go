package main

import (
	"os"
	"sync"
)

type Mutex struct {
	ch chan int
}

func NewMutex() *Mutex {
	mu := &Mutex{
		ch: make(chan int, 1),
	}
	mu.ch <- 1
	return mu
}

func (m Mutex) Lock() {
	<-m.ch
}

func (m Mutex) UnLock() {
	m.ch <- 1
}

func main() {
	file, _ := os.Create("test.txt")
	defer func() {
		file.Close()
	}()

	mu := NewMutex()

	list := []string{"tony", "sherry", "tim", "dom"}

	var wg sync.WaitGroup
	wg.Add(len(list))

	for _, item := range list {
		go func(name string) {
			mu.Lock()
			for i := 0; i < 10000; i++ {
				file.WriteString(name + "\n")
			}
			mu.UnLock()
			wg.Done()
		}(item)
	}

	wg.Wait()
}
