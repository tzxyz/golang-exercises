package main

import "fmt"

type Integer int

/**
 *	go可以根据func (i Integer) Less(b Integer) bool
 *	自动生成  func (i *Integer) Less(b Integer) bool
 *	但是不会根据func (i *Integer) Add(b Integer)
 *	自动生成    func (i Integer) Add(b Integer)
 */
type LessAdder interface {
	Less(b Integer) bool
	Add(b Integer)
}

func (i Integer) Less(b Integer) bool {
	return i < b
}

func (i *Integer) Add(b Integer){
	*i += b
}

func main() {
	/*var i LessAdder = new(Integer)
	fmt.Println(i.Less(8))*/

	var i Integer = 10
	var les LessAdder = &i
	fmt.Println(les.Less(10))
	fmt.Println(les.Less(20))
	// var les LessAdder = i
	// 为什么这句话不能对接口赋值
}
