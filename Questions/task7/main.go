package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v

			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
		close(c)
	}()

	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)

	go func() {
		for {
			select {
			case v := <-a: // как только канал опустеет будут слаться пустые значения, для int это 0
				c <- v
			case v := <-b:
				c <- v
			}
		} // нужно правильно закрыть канал С
	}()

	return c
}

// рабочий merge
func or(channels ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(channels))
	for _, c := range channels {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4, 6, 8)
	c := or(a, b)

	for v := range c {
		fmt.Println(v)
	}
}
