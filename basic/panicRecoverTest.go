package main

import (
	"fmt"
	"time"
)

/*
跨协程失效
panic 只会触发当前 Goroutine 的延迟函数调用
*/
func boundInG() {
	defer println("in main")
	go func() {
		defer println("in goroutine")
		panic("")
	}()
	time.Sleep(1 * time.Second)
}

/*
崩溃恢复
recover 只有在发生 panic 之后调用才会生效。
需要在 defer 中使用 recover 关键字。
*/
func panicRecover() {
	defer fmt.Println("in main")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("err msg: ", r)
		}
	}()
	//if err := recover(); err != nil {
	//	fmt.Println(err)
	//}
	panic("unknown err")
}

/*
嵌套崩溃
多次调用 panic 也不会影响 defer 函数的正常执行。所以使用 defer 进行收尾的工作一般来说都是安全的。
*/
func panicInLoop() {
	defer fmt.Println("in main")
	defer func() {
		defer func() {
			panic("panic again and again")
		}()
		panic("panic again")
	}()
	panic("panic once")
}

func main() {
	//boundInG()

	//panicRecover()

	panicInLoop()
}
