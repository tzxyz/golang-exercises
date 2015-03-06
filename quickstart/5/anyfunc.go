package main

import "fmt"

func main() {
	/**
		匿名函数：
		没有函数名的函数
		可以直接赋值给变量
		返回值只有一个的话不许要加小括号
	*/
	f := func(a int, b int) int {
		return a + b
	}
	fmt.Println(f(1, 3))

	/**
		匿名函数后面加上参数列表表示直接调用
	 */
	var text string = func(args ...string) string {
		return args[0] + args[1] + args[2]
	}("who", "what", "why")

	fmt.Println(text)
}
