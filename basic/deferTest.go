package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i) //4 3 2 1 0
	}

	/*
		预计算参数
		调用 defer 关键字会立刻对函数中引用的外部参数进行拷贝，所以 time.Since(startedAt) 的结果不是在 main 函数退出之前计算的，
		而是在 defer 关键字调用时计算的，最终导致上述代码输出 0s。
	*/
	startedAt := time.Now()
	defer fmt.Println(time.Since(startedAt)) //0s
	time.Sleep(time.Second)

	/*
		虽然调用 defer 关键字时也使用值传递，但是因为拷贝的是函数指针，所以 time.Since(startedAt) 会在 main 函数执行前被调用并打印出符合预期的结果。
	*/
	startAt := time.Now()
	defer func() { fmt.Println(time.Since(startAt)) }()
	time.Sleep(time.Second)
}
