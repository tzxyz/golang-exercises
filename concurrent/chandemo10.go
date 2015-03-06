package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

// 简单的工厂模式分发
// 返回一个带数据的channel
func build() chan int {
	c := make(chan int)
	rand.Seed(time.Now().UnixNano())

	go func() {
		time.Sleep(time.Second)
		// 为什么rand.Int()这么大
		i := rand.Int()
		fmt.Println(i)
		c <- i
	}()
	return c
}

func main() {
	c := build()
	fmt.Println("random int = ", <-c)
}
