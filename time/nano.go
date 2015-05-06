package main

import (
	"fmt"
	"time"
)

func main() {
	// println(NowTime())
	// println("log ----->>> ")
	DatePattern()
	// Ticker()
}

//返回当前时间毫秒值
func NowTime() int64 {
	now := time.Now()
	nt := now.UnixNano()
	return nt
}

func Ticker() {
	ticker := time.NewTicker(time.Second) //创建一个定时器
	go tick(ticker.C)
	time.Sleep(1 * time.Second)
	<-time.After(1 * time.Second) //1秒后停止定时器
	ticker.Stop()
}

func DatePattern() {
	t, err := time.Parse("2006-01-02 15:04:05", "2015-01-12 20:22:55")
	fmt.Println(t)
	fmt.Println(err)
}

func tick(ch <-chan time.Time) {
	for t := range ch {
		fmt.Println(t)
	}
}
