package main

import "fmt"
import "reflect"

type User struct{
	Uid string
	age int
}

/**
 *	通过反射修改对象的值
 *	1.获取Value
 *	2.判断v的类型是否是指针(必须是指针) && 判断v是否可以被修改
 *	3.找出v要修改的字段(该字段必须存在)，使用对应的方法修改
 */
func SetUser(o interface{}){
	v:=reflect.ValueOf(o)
	if !(v.Kind()==reflect.Ptr && v.Elem().CanSet()) {
		fmt.Println("non-canset")
	}else{
		v = v.Elem()
	}

	if f:=v.FieldByName("Uid");f.IsValid() && f.Kind() == reflect.String {
		f.SetString("uid009")
	}
}

func main() {
	var u User = User{Uid:"uid001",age:22}
	fmt.Println("修改前",u)
	SetUser(&u)
	fmt.Println("修改后",u)
}
