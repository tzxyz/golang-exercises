package main

import "fmt"

func main() {
	/**
	 *	1.defer会等函数体执行完毕后才执行
	 *	2.函数体内的变量做为defer的匿名函数的参数时
	 *	  在定义defer定义时，就获取了值的拷贝
	 *	3.函数体内的变量不做为defer匿名函数的参数时
	 *	  引用的则是变量的内存地址
	 *	4.defer定义的函数必须被调用
	 */
	for i := 0; i < 5; i++ {
		defer func(i int) {
			fmt.Printf("%d ", i)
			i++
		}(i)
		fmt.Println(i)
	}

	//	可以看出defer匿名函数的变量(非做为参数传递)
	//  中的引用的是函数体的内存地址
	for i := 0; i < 5; i++ {
		defer func() {
			fmt.Printf("%p ",&i)
		}()
		fmt.Printf("%p ",&i)
	}
}
