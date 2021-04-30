package main

import (
	"fmt"
	"time"
)

func Fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
			fmt.Println(x, y)
		case <-quit:
			fmt.Println("quit...")
			return
		default: //无default 语句时，会阻塞当前的 Goroutine
			fmt.Println("default...")
			return
		}
	}
}

/*
随机执行,多个case条件同时满足，随机选择一个执行，引入随机性来避免饥饿
*/
func RandomExec() {
	ch := make(chan int)
	go func() {
		for range time.Tick(1 * time.Second) {
			ch <- 0
		}
	}()
	for {
		select {
		case <-ch:
			println("case1")
		case <-ch:
			println("case2")
		}
	}
}

func main() {
	c := make(chan int, 5)
	q := make(chan int)
	Fibonacci(c, q)

	RandomExec()
}
