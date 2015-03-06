package main

import (
	"fmt"
	"runtime"
	"time"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	names := []string{"卓尼玛", "陈尼玛", "王尼玛"}

	// 这样写容易角标越界
	// 比如第一次循环的时候i=0,go func()还没执行,i已经加到1了,这时候打印出来的就不对
	// for i := 0; i < len(names); i++ {
	// 	go func() {
	// 		fmt.Println(names[i])
	// 	}()
	// }

	// 这样子写打印出来的有可能都是最后一个
	// goroutine的调度不可控制
	// 原因是name被赋值后并没有马上被执行
	for _, name := range names {
		go func() {
			fmt.Println(name)
		}()
	}

	runtime.Gosched()
	time.Sleep(time.Second)

	// 应该避免以上情况
	for i := 0; i < len(names); i++ {
		go func(name string) {
			fmt.Printf("hi,%s\n", name)
		}(names[i])
	}

	runtime.Gosched()
	time.Sleep(time.Second)
}
