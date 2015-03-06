package main

import (
	"fmt"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

// chan是引用类型
// 默认是同步模式,需要发送和接收配对,否则会被阻塞

func main() {
	data := make(chan int)
	exit := make(chan bool)

	go func() {
		// 从data中遍历数据，直到这个channel被close()
		for d := range data {
			fmt.Println(d)
		}
		fmt.Println("Receive Done")
		// 发送退出通知
		exit <- true
	}()

	go func() {
		for i := 0; i < 100; i++ {
			data <- i
		}
		// 为什么这句话注释了也会产生死锁?
		close(data)
		fmt.Println("Send Done")
	}()

	// 等待退出通知
	<-exit
}
