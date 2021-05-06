package main

import (
	"context"
	"fmt"
	"time"
)

//多个 Goroutine 同时订阅 ctx.Done() 管道中的消息，一旦接收到取消信号就停止当前正在执行的工作并提前返回。
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go handle(ctx, 1500*time.Millisecond)

	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
	}
}
func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())
	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}
}
