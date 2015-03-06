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
	name := "卓尼玛"
	// 为什么这里name没有变化呢,因为name是我传递进去的
	// go func(name string) {
	// 	fmt.Printf("my name is %s", name)
	// Output:"my name is zhuonima"
	// 原因runtime.Gosched()之前他已经执行了
	// }(name)
	go func() {
		fmt.Println(name)
	}()
	// runtime.Gosched()
	name = "王尼玛"
	runtime.Gosched()
	time.Sleep(time.Second)
}
