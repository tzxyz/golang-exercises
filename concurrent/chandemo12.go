package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	// 为什么这里使用var声明之后wg就可以直接调用函数了
	var wg sync.WaitGroup

	// 同步的channel
	quit := make(chan bool)

	// 循环往wg中添加goroutine
	for i := 0; i < 3; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			task := func() {
				fmt.Println(id, time.Now().Nanosecond())
				time.Sleep(2 * time.Second)
			}
			// 随机选择一个goroutine
			// 如果是quit,整个goroutine退出
			// 否则运行task
			for {
				select {
				// 当quit被关闭的时候,这里就不会阻塞了
				case <-quit:
					return
				default:
					task()
				}
			}
		}(i)
	}

	// 尝试让程序运行10秒,然后关闭quit
	time.Sleep(10 * time.Second)
	close(quit)
	// 等待wg中所有的goroutine结束
	wg.Wait()
}
