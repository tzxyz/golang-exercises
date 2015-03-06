package main

import "fmt"

type Integer int

/**
以下两种方式的区别？
指针并没有改变a的值
*/
func (a Integer) add(b Integer) Integer {
	a += 1
	return a + b
}

func (a *Integer) add2(b Integer) Integer {
	*a += 1
	return *a + b
}

func (a Integer) Add(b Integer) {
	a += b
}

func (a *Integer) Add2(b Integer) {
	*a += b
}

func main() {
	var a Integer = 1
	fmt.Println(a.add(1))
	fmt.Println(a)
	fmt.Println(a.add(2))
	fmt.Println(a)

	a = 2
	var b Integer = 1
	fmt.Println(a.Add(b))
	fmt.Println(a.Add2(b))
}
