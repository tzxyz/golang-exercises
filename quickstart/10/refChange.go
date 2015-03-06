package main

import "fmt"
import "reflect"

func main() {
	x := 5
	v := reflect.ValueOf(&x)
	fmt.Println(x)
	v.Elem().SetInt(999) //通过反射改变x的值
	fmt.Println(x)

	f := 3.14
	v = reflect.ValueOf(&f)
	fmt.Println(f)
	v.Elem().SetFloat(9527.77)
	fmt.Println(f)
}
