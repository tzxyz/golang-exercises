package main

import (
	"container/ring"
	"fmt"
)

func DoDemo() {
	r := ring.New(10)
	// 这样子并没有赋值,因为go中都是值拷贝
	// r.Do(func(v interface{}) {
	// 	v = 111
	// })
	for i := 0; i < r.Len(); i++ {
		r.Value = i
		r = r.Next()
	}
	r.Do(func(v interface{}) {
		fmt.Println(v)
	})
}

func NewDemo() {
	// 创建一个回环链表
	r := ring.New(10)
	fmt.Println("len(r)", r.Len())
}

func main() {
	// NewDemo()
	DoDemo()
}
