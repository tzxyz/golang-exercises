package main

import (
	"fmt"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

// 异步的channel
// 异步的channel存在缓冲区
// 当缓冲区满的时候,发送被阻塞,当缓冲区空的时候,接收被阻塞
func main() {
	data := make(chan int, 3)
	exit := make(chan bool)

	go func() {
		for d := range data {
			// 使用内置函数cap()查看缓冲区的容量
			// 使用内置函数len()查看缓冲区元素个数
			fmt.Println("cap=", cap(data))
			fmt.Println("len=", len(data))
			fmt.Println("data=", d)
		}
		// 发送退出通知
		exit <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			data <- i
		}
		// 异步的channel同样需要关闭
		// 没关闭会造成死锁
		close(data)
	}()

	<-exit
}
