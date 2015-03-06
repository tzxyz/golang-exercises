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
	ch := make(chan int, 5)
	sign := make(chan byte, 2)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
		fmt.Println("ch close")
		sign <- 0
	}()

	go func() {
		for {
			v, ok := <-ch
			if !ok {
				break
			}
			fmt.Println("v=", v)
		}
		sign <- 1
	}()

	// 这里为什么不会死锁?发送端并没关闭
	// 缓冲区空了 为什么没有阻塞
	<-sign
	<-sign
}
