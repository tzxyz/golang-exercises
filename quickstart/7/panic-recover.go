package main

import "fmt"

func A() {
	fmt.Println("func a")
}

/**
 *	函数体b执行到panic，出现错误
 *	不执行panic之后的内容
 *	函数体执行完毕进入defer处理recover
 */	
func B() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("defer func b")
		}
	}()
	panic("panic in b")
	fmt.Println("func b")
}

func C() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("defer func c")
		}
	}()
	panic("panic in c")
}

func main() {
	A()
	B()
	C()
}
