package main

import (
	"fmt"
	"runtime"
	"sync"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

// runtime.Goexit()立即终止当前goroutine
// 调度器确保所有被注册的defer被逆序执行
// 为什么会死锁！！！
func main() {
	wg := new(sync.WaitGroup)
	wg.Add(runtime.NumCPU())

	go func() {
		defer wg.Done()
		defer fmt.Println("defer (-.-) ")
		func() {
			defer fmt.Println("defer next")
			runtime.Goexit()
			println("next")
		}()
		println("once")
	}()
	wg.Wait()
}
