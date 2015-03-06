package main

import (
	"fmt"
	"reflect"
)

/**
 *	私有字段不能使用TypeOf和ValueOf反射？
 */
type User struct {
	Name        string
	Age         int
	Description string
}

/**
 *	私有函数确可以反射？
 */
func (u User) hello(){
	fmt.Println("my name is ",u.Name,", im ",u.Age," years old,",u.Description)
}

/**
 *	TypeOf反射出对象的struct
 *	ValueOf反射出对象的值
 */
func Info(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("type :", t.Name())
	v := reflect.ValueOf(o)
	fmt.Println("value :", v)

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)		//获取出第i个字段
		val := v.Field(i).Interface()	//获取第i个字段的值
		fmt.Println(f.Name,"=",val,":",f.Type)
	}
}

func CallMethod(o interface{}){
	t := reflect.TypeOf(o)
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Println(m)
		fmt.Println(m.Name,m.Type)
	}
}

func main() {
	u := User{Name: "卓尼玛", Age: 22, Description: "im gopher！！！"}
	Info(u)
	fmt.Println()
	CallMethod(u)
}
