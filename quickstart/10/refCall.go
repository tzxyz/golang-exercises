package main

import "fmt"
import "reflect"

type User struct {
	Uid  string
	Name string
}

func (u User) SayHello(name string) {
	fmt.Println("hello,", name, ",my name is ", u.Name)
}

func main() {
	u := User{Uid: "uid001", Name: "卓尼玛"}
	v := reflect.ValueOf(u)
	m := v.MethodByName("SayHello")                        //1.反射要调用的方法
	args := []reflect.Value{reflect.ValueOf("xiao bitch")} //2.这是一个切片
	m.Call(args)                                           //3.调用方法
}
