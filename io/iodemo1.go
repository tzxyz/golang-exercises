package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// CopyDemo()
	// CopyNDemo()
	// WriteStringDemo()
	// ReadFileDemo()
	// ReadDirDemo()
	// ReadAllDemo()
	// TempFileDemo()
	TempDirDemo()
}

func TempDirDemo() {
	// 在dir下创建一个前缀名为prefix的临时文件夹
	// 如果dir为空字符串,则表示系统默认的临时文件目录(os.TempDir)
	td, err := ioutil.TempDir("", "tempd")
	if err != nil {
		log.Fatal("create temp dir error")
	}
	fmt.Println(td)
}

func TempFileDemo() {
	// 在dir目录下创建一个前缀为tmp临时文件
	// 如果dir是空字符串,则表示当前系统默认的临时文件目录(os.TempDir)
	// 可以用于图片上传回显
	tf, err := ioutil.TempFile("", "tmp")
	if err != nil {
		log.Fatal("create temp file error")
	}
	fmt.Println(tf.Name())
	fmt.Println(os.TempDir())
}

func ReadAllDemo() {
	r, _ := os.Open("iodemo1.go")
	// 将一个流中的数据全部读取出来
	bytes, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bytes))
}

func ReadDirDemo() {
	// 读取一个目录,获取该目录下的所有fileInfo
	files, err := ioutil.ReadDir("/Users")
	if err != nil {
		log.Fatal(err)
	}
	for i, file := range files {
		fmt.Println(i, "=", file.Name())
	}
}

func ReadFileDemo() {
	// 读取一个文件,返回总共读取的字节数
	bytes, err := ioutil.ReadFile("iodemo1.go")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bytes))
}

func WriteStringDemo() {
	w, _ := os.Create("write-demo.go")
	// 并不是追加,而是覆盖
	n, err := io.WriteString(w, "fuck my life")
	if err != nil {
		log.Fatal(err)
	}
	println("writtern in", n)
}

func CopyNDemo() {
	src, _ := os.Open("nano.go")
	dst, _ := os.Create("nano-copyn.go")
	writtern, err := io.CopyN(dst, src, 200)
	// 表示拷贝成功了
	if err == nil {
		println(writtern)
		println("拷贝成功")
	}
	// 如果n比src的总bytes数还大,则会出现这个异常
	if err == io.EOF {
		println(writtern)
		println("字数不够拷贝啦")
	}
	// 什么情况下写入数据会不等于读取数据 ###
}

func CopyDemo() {
	// 打开一个文件
	src, _ := os.Open("nano.go")
	// 创建一个文件
	dst, _ := os.Create("nano-copy.go")
	// 将src流中的数据复制到dst中,返回的是总共复制的字节数
	writtern, err := io.Copy(dst, src)
	if err != nil {
		log.Fatal("copy err ...", err)
	}
	println("writtern bytes", writtern)
}
