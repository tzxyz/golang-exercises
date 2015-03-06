package main

import (
	"fmt"
)

//struct必须要写花括号
type File struct{
	filename string
	size int
}

/**
 *	非侵入式接口：
 *	Ifile接口定义了4个方法
 *	File类只要实现了这4个方法，就是实现了该接口
 *	方法签名要和接口的一致
 */
type Ifile interface {
	Read(buf []byte) (n int, err error)
	Write(buf []byte) (n int, err error)
	Seek(off int, whence int) (pos int, err error)
	Close() error
}

func (f *File) Read(buf []byte) (n int, err error){
	return 0,nil
}
func (f *File) Write(buf []byte) (n int, err error){
	return 0,nil
}
func (f *File) Seek(off int, whence int) (pos int, err error){
	return 0,nil
}
func (f *File) Close() error{
	return nil
}

func main() {
	//只要File类实现了接口，就可以进行对接口赋值
	var file Ifile = new(File)
	fmt.Println(file)
}
