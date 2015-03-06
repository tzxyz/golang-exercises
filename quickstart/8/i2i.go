package main

import "fmt"

type File struct{}

type ReaderAndWriter interface {
	Read()
	Write()
}

type InputStreamAndOutputStream interface {
	Write()
	Read()
}

func (f *File) Read() {

}

func (f *File) Write() {

}

/**
 *	1.接口赋值并不要求两个接口必须等价。
 *	2.如果接口A的方法列表是接口B的方法列表的子集，
 *	  那么接口B可以赋值给接口A
 */
func main() {
	// 1.将f1对象赋值给接口
	var f1 ReaderAndWriter = new(File)
	// 2.将接口赋值给接口
	var f2 InputStreamAndOutputStream = f1
	var f3 ReaderAndWriter = f2

	fmt.Println(f1 == f2)
	fmt.Println(f2 == f3)
}
