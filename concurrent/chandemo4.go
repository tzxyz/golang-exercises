package main

import (
	"fmt"
	"runtime"
	"sync"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

// runtime.Gosched()类型线程礼让
// 当前goroutine放弃对CPU的执行权
// 注意go func()的时候必须wg.Done(),否则会造成死锁
// 为什么会造成死锁呢?
func main() {
	wg := new(sync.WaitGroup)
	wg.Add(runtime.NumCPU())
	for i := 0; i < runtime.NumCPU(); i++ {
		go func(id int) {
			defer wg.Done()
			for i := 0; i < 10; i++ {
				fmt.Printf("id = %d,count = %d\n", id, i)
			}
			runtime.Gosched()
		}(i)
	}
	defer wg.Wait()
}
