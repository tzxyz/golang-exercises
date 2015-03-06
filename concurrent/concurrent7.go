package main

import (
	"fmt"
)

func sum(s []int, ch chan int) {
	sum := 0
	for _, n := range s {
		sum += n
	}
	ch <- sum
}

func main() {
	var s [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ch1 := make(chan int)			//创建两个通道
	ch2 := make(chan int)
	go sum(s[:len(s)/2], ch1)		//开启一个goutine计算数组的前一半的总和
	go sum(s[len(s)/2:], ch2)

	x, y := <-ch1, <-ch2			//分别用x，y接收用通道接收计算好的sum
	fmt.Printf("x=%d,y=%d", x, y)
}
