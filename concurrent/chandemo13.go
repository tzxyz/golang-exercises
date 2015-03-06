package main

import (
	"fmt"
	"runtime"
	"time"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	w := make(chan bool)
	c := make(chan int)

	go func() {
		select {
		case v := <-c:
			fmt.Println("v=", v)
		case <-time.After(3 * time.Second):
			fmt.Println("timeout")
		}
		// 这里没有发送w就阻塞了,死锁
		w <- true
	}()
	// c <- 1	注释掉以后-<c阻塞,引发timeout
	<-w
}
