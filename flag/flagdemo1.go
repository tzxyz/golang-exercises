package main

import (
	"flag"
)

func main() {
	// ArgDemo()
	// ArgsDemo()
}

func ArgDemo() {
	// 解析命令行
	flag.Parse()
	// 返回一个string类型的值,命令行的第i+1个参数的值
	// 使用前必须先解析命令行
	arg1 := flag.Arg(0)
	arg2 := flag.Arg(1)
	arg3 := flag.Arg(2)
	// go run flagdemo1.go 111 222 333
	println(arg1)
	println(arg2)
	println(arg3)
}

func ArgsDemo() {
	flag.Parse()
	// 获取命令行的所有参数
	args := flag.Args()
	for i, arg := range args {
		println("args[", i, "]", arg)
	}
}
