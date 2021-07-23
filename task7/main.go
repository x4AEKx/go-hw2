package main

import (
	"fmt"
	"math/rand"
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

func main() {
	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4, 6, 8)
	c := merge(a, b)

	for v := range c {
		fmt.Println(v)
	}
}
