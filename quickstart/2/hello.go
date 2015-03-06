package main
// 1.可执行的必须是main包

import "fmt"
// 2.导入的包

const PI = 3.14
// 3.常量的声明

var name = "gopher"
// 4.全局变量的声明

type number int
// 5.一般类型变量的声明

type gohper struct{}
// 6.结构类型的变量声明

type golang interface{}
// 7.接口的声明

func main() {
	fmt.Println("hello world 你好，世界");
}
// 8.函数的声明