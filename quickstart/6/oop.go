package main

import "fmt"

type Integer int

/**
	func 后面的(a Integer)是什么
	两个Integer类型是怎么比较的
 */
func (a Integer) less(b Integer) bool {
	return a < b
}
func main() {
	var a Integer = 1
	flag := a.less(2)
	fmt.Println(flag)
}
