package main

import (
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

// 死锁现象
// 将一个没有缓冲区的channel分成两个单向channel
// send端发送后就死锁
// 因为单向channel没有人接收
// 需要设置缓冲区
// *不能将单向channel转换成普通channel
func main() {
	// 将send声明成单向的channel(发送)
	// 将recv声明成单向的channel(接受)
	c := make(chan int, 3)
	var send chan<- int = c
	var recv <-chan int = c

	send <- 1
	send <- 2
	send <- 3
	// 单向的channel不能发送比缓冲区容量还大的个数
	// 因为当缓冲区满的时候,发送端继续发送会导致阻塞
	// send <- 4

	println(<-recv)
	println(<-recv)
	println(<-recv)

	// 同理当缓冲区空的时候,接收端继续接收也会导致阻塞
	println(<-recv)

	// 只接收的channel是不能关闭的
	close(send)
	// close(recv)
}
