package main

import (
	"fmt"
	"time"
)

func main() {
	// AfterFuncDemo()
	// NewTimerDemo()
	// TimerStopDemo()
	// AddDemo()
	SubDemo()
}

func SubDemo() {
	now := time.Now()
	time.Sleep(5 * time.Second)
	after := time.Now()
	sub := after.Sub(now) //两个时间相减 和Add函数不同 它的参数是time.Time
	fmt.Println(sub)
}

func AddDemo() {
	now := time.Now()
	fmt.Println(now)
	after := now.Add(time.Hour) //时间相加 参数是time.Duration
	fmt.Println(after)
}

func TimerStopDemo() {
	b := time.AfterFunc(time.Second, func() { //停止timer 如果这个timer已经结束或过期 返回false
		fmt.Println("callback")
	}).Stop() //还没执行就关闭了
	fmt.Println(b)
	timer := time.AfterFunc(time.Second, func() {
		fmt.Println("callbacl next")
	})
	time.Sleep(5 * time.Second)
	b = timer.Stop()
	fmt.Println(b) //已经结束了 返回false
}

func NewTimerDemo() {
	timer := time.NewTimer(time.Second * 5) //创建一个定时器 5秒后将当前时间发送给C
	t := <-timer.C
	fmt.Println(t)
}

func AfterFuncDemo() {
	f := func() {
		fmt.Println("this is a timer")
	}
	time.AfterFunc(time.Second, f) //经过时间d后,在自己的goroutine中执行函数f
	time.Sleep(time.Second * 5)
}
