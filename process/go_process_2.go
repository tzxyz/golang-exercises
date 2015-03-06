package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])
}

// 1.进程名称
// 进程名称一般都是进程参数的第一个字符串
// 使用go run 每次都会重新编译,运行程序,所以每次进程的名称不同
// go build编译之后则不会,编译之后的名称都是固定的
