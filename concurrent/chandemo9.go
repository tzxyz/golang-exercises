package main

import (
	"fmt"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

// 疑问:使用select必须要开启缓冲区吗
// select和switch类似,只执行一次，随机选择一个可用的channel
// 循环中使用select default case要小心，避免形成洪水(什么意思?)
func main() {
	a := make(chan int, 5)
	b := make(chan bool, 5)

	go func() {
		for {
			// 随机选择一个channel发送数据
			select {
			case a <- 1:
				fmt.Println("a send")
			case b <- true:
				fmt.Println("b send")
			}
		}
	}()

	for {
		// 随机选择一个channel,读取里面的数据
		select {
		case v, _ := <-a:
			fmt.Println("a revc ", v)
		case v, _ := <-b:
			fmt.Println("b revc ", v)
		}
	}

	//用于阻塞main goroutine
	select {}
}
