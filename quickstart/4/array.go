package main

import "fmt"

func main() {
	a := [...]string{"who", "what", "why"}
	p := &a
	fmt.Println(p)

	//指针上面不能使用[...]
	b := [...]string{"son", "of", "bitch"}
	var q *[3]string = &b
	fmt.Println(q)

	//数组里面是什么类型，就要定义什么类型的数组
	x, y := 1, 2
	c := [...]*int{&x, &y}
	fmt.Println(c)
}
