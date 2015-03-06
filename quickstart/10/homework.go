package main

import "fmt"
import "reflect"

type Person struct {
	Pid         string
	Name        string
	Age         int
	Description string
}

func (p Person) Hard2Work(job string) {
	fmt.Println("my job:", job)
}

func PrintfInfo(o interface{}) {
	t := reflect.TypeOf(o)
	v := reflect.ValueOf(o)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Println(f.Name, ":", val)
	}
}

func CallMethod(o interface{}, methodName, job string) {
	v := reflect.ValueOf(o)                       //1.反射出value
	m := v.MethodByName(methodName)               //2.反射方法
	args := []reflect.Value{reflect.ValueOf(job)} //3.调用方法需要的参数
	m.Call(args)                                  //4.方法调用
}

func main() {
	var p Person = Person{Pid: "pid001", Name: "卓尼玛", Age: 22, Description: "年薪百万！"}
	PrintfInfo(p)
	CallMethod(p, "Hard2Work", "Programmer")
}
