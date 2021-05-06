package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"
	"golang.org/x/sync/errgroup"
)
var g errgroup.Group

func main() {
	onceTest()
	condTest()
}

func waitGroupTest() {
	requests := []*Request{}
	wg := &sync.WaitGroup{}
	wg.Add(len(requests))
	for _, request := range requests {
		go func(r *Request) {
			defer wg.Done()
			// res, err := service.call(r)
		}(request)
	}
	wg.Wait()
}

type Request struct {
}

func noCopyTest() {
	wg := sync.Mutex{}
	yawg := wg
	fmt.Println(wg, yawg)
}

func onceTest() {
	o := &sync.Once{}
	for i := 0; i < 10; i++ {
		o.Do(func() {
			fmt.Println("only once")
		})
		fmt.Println("outer loop...", i)
	}
}

func condTest() {
	c := sync.NewCond(&sync.Mutex{})
	for i := 0; i < 10; i++ {
		go listen(c)
	}
	time.Sleep(1 * time.Second)
	go broadcast(c)
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
}

func broadcast(c *sync.Cond) {
	c.L.Lock()
	c.Broadcast() //广播唤醒等待队列中的协程，按入队顺序唤醒
	c.L.Unlock()
}
func listen(c *sync.Cond) {
	c.L.Lock()
	c.Wait() //阻塞，并使得协程休眠入链表等待
	fmt.Println("listen")
	c.L.Unlock()
}


func errGroupTest() {
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
	}
	for i := range urls {
		url := urls[i]
		g.Go(func() error {
			resp, err := http.Get(url)
			if err == nil {
				resp.Body.Close()
			}
			return err
		})
	}
	if err := g.Wait(); err == nil {
		fmt.Println("Successfully fetched all URLs.")
	}
}
