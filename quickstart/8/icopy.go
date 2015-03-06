package main

import "fmt"

type Say interface {
	SayHello()
}

type Man struct {
	name string
}

func (m Man) SayHello() {
	fmt.Println("my name is ", m.name)
}

func main() {
	var m Man = Man{name: "卓尼玛"}
	var s Say = m

	/**
	 *	对象赋值给接口时，接口会拷贝一份复制品
	 *	接口引用的是这个复制品的指针
	 *	源对象无法改变接口复制品的状态，也无法改变指针
	 */
	s.SayHello()
	m.name = "fuck my life"
	// s.name = "fuck my life"	接口里面没有name属性
	s.SayHello()
}
