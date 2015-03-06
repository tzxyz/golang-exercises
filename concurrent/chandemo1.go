package main

import (
	"fmt"
	"runtime"
	"time"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

// 当没有开启runtime.GOMAXPROCS(runtime.NumCPU())的时候是按顺序执行的
// 没有开启runtime时候执行顺序是 main - fmt1 - fmt2
// 为什么?
// 因为进程启动后默认只允许一个系统线程服务与goroutine
// 可以通过环境变量来配置或者使用runtime包的函数
// 调度器无法控制goroutine的执行顺序
// 把runtime.GOMAXPROCS(runtime.NumCPU())放到init里面好点
func main() {
	go func() {
		for i := 0; i < 30; i++ {
			fmt.Println("go fmt1")
		}
	}()
	go func() {
		for i := 0; i < 30; i++ {
			fmt.Println("go fmt2")
		}
	}()
	for i := 0; i < 30; i++ {
		fmt.Println("main")
	}
	time.Sleep(time.Second)
}
