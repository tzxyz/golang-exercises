package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1:])
}

// 1.进程参数
// 任何进程启动时,都可以赋予一个字符串数组作为参数
// os.Args的第一个参数是进程名称
