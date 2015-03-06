package main

import (
	"fmt"
	"runtime"
	"sync"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(3)

	sem := make(chan int, 1)

	for i := 0; i < 3; i++ {
		go func(id int) {
			defer wg.Done()
			// 向自己所在的goroutine发送数据
			// 如果这个sem里面的缓冲区已经有数据了,那么就阻塞了
			// 如果一个goroutine发送了数据,必须等这个goroutine接收了数据其他goroutine才能执行
			sem <- 1
			for i := 0; i < 3; i++ {
				fmt.Println(id, i)
			}
			// 接收自己所在的goroutine的数据，如果阻塞其他的goroutine可以执行?
			// 不能执行
			<-sem
		}(i)
	}
	wg.Wait()
}
