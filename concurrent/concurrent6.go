package main

import (
	"fmt"
	"time"
)

func main() { //select的超时机制：只有其中一个case执行完成，程序就会继续执行
	ch := make(chan int)
	timeout := make(chan bool)
	go func() {
		time.Sleep(1e9)
		timeout <- true
	}()
	for {
		select {
		case i := <-ch:
			fmt.Println(i)
		case <-timeout:
		}
		ch <- 1 //这里为什么会死锁
	}
}
