package main

import (
	"flag"
	"fmt"
)

func main() {
	wordPtr := flag.String("word", "foo", "a string")
	// 参数名,默认值,描述

	numPtr := flag.Int("num", 42, "a num")

	boolPtr := flag.Bool("form", true, "a bool")

	var str string

	flag.StringVar(&str, "svar", "a strvar", "a strvar")
	//指定了一个指针 跟flag.String()差不多

	flag.Parse()

	fmt.Println("word:", *wordPtr) //对一个指针取地址?
	fmt.Println("num", *numPtr)
	fmt.Println("bool", *boolPtr)
	fmt.Println("todo", flag.Args())

	//这篇还是进程参数
	//编译之后 ./go_process_4 -help
	//用法和linux命令差不多
	//question:flag.Args()和os.Args()应该是一样的
}
