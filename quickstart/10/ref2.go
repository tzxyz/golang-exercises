package main

import "fmt"
import "reflect"

type User struct {
	Id   string
	Name string
}

type Manager struct {
	User
	Title string
}

func showInfo(o interface{}) {
	t := reflect.TypeOf(o)
	k := t.Kind() //判断对象是什么类型的
	// fmt.Println(k)

	//	如果不是结构体就返回
	if k != reflect.Struct {
		fmt.Println("non-struct")
		return
	}

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		v := reflect.ValueOf(o).Field(i).Interface()
		fmt.Println(f.Name, ":", v, ":", f.Type)
	}
}

func main() {
	var m Manager = Manager{User: User{"uid001", "卓尼玛"}, Title: "超级管理员"}
	showInfo(m)
	showInfo(1)

	t :=reflect.TypeOf(m)

	// 取出User中的字段：
	// 第一个参数0表示Manager中的第一个Field
	// 第二个参数0表示User中的第一个Field
	fmt.Println("%#v\n",t.FieldByIndex([]int{0,0}))
	fmt.Println("%#v\n",t.FieldByIndex([]int{0,1}))
}
