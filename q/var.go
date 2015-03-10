package main

import (
	"fmt"
	"reflect"
)

type A struct {
	a string
	b int64
	c *A
}

func (this *A) Print() {
	fmt.Println("a")
}

func main() {
	var a1 A
	var a2 *A

	fmt.Println(a1)
	fmt.Println(a2)

	a1.Print()
	a2.Print() // nil为什么能调用Print()

	fmt.Println(reflect.TypeOf(a1))
	fmt.Println(reflect.TypeOf(a2))

}
