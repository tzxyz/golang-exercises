package main

import (
	"fmt"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

// 产生了死锁
// 创建了一个匿名通道，只有接受端没有发送端
func main() {
	<-(make(chan bool))
	fmt.Println("Done")
}
