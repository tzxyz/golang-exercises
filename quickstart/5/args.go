package main

import (
	"fmt"
)

/*
	go同样支持可变参数
	可以传入数组，也可以传入切片
	测试是不能传递数组和切片的
*/
func myfunc1(args ...string) {
	for c, str := range args {
		fmt.Println(c, str)
	}
}

/**
	args... interface{}用于接受任意类型的参数
 */
func myfunc2(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is a int value")
		case string:
			fmt.Println(arg, "is a string value")
		default:
			fmt.Println(arg, "is another value")
		}
	}
}

func main() {
	myfunc1("im vincent")
	myfunc1("who", "what", "why")

	/*arr := [...]string{"here", "there"}
	slice := arr[:1]
	myfunc1(arr)
	myfunc1(slice)*/

	var v1 int = 15
	var v2 string = "fuck my life"
	var v3 float64 = 3.141592658
	myfunc2(v1, v2, v3)
}
