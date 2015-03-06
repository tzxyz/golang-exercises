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

// 互斥锁sync.Mutex
// 一个互斥锁可以锁不同的goroutine
// 互斥锁被锁定了,其他goroutine就不能在对他上锁了
// 可以看到如果main没有Unlock(),for循环中的就没办法执行Lock()之后的语句
// question:为什么有时候会打印主goroutine还没unlock就有goroutinelock了?
// 因为第37行unlock之后,调度器不一定会继续让主goroutine执行,有可能其他goroutine执行了
// 可以看到的现象是只要主goroutine没有解锁,其他的goroutine就不会上锁

// 将time.Sleep(time.Secound)替换成runtime.Gosched()现象是一样的
// 只是有时候主goroutine拿到执行权后结束就不打印了
func main() {
	var mux sync.Mutex
	mux.Lock()
	// 3.如果是defer解锁呢,会出现只有主goroutine打印
	// 因为defer是在函数执行返回之前调用,这时候主函数结束了才解锁,其他的goroutine就结束了,所以不打印
	// defer mux.Unlock()
	fmt.Println("Lock is locked. (Go)")
	for i := 1; i <= 3; i++ {
		go func(id int) {
			// 2.没有这句话,只有一个goroutine打印Lock is locked
			// 因为这个goroutine没有解锁,其他的goroutine就不能上锁
			defer mux.Unlock()
			mux.Lock()
			fmt.Printf("Lock is locked. (Go%d)\n", id)
		}(i)
	}
	// runtime.Gosched()
	time.Sleep(time.Second)
	// 1.没有mux.Unlock(),其他的goroutine就没办法进行上锁操作
	mux.Unlock()
	fmt.Println("Lock is unlocked. (Go)")
	time.Sleep(time.Second)
	// runtime.Gosched()
}
