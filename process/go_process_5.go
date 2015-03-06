package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	resp, _ := http.Get("http://www.huaban.com")
	// resp.Body
	html, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(html))
	StandardIn()
	StandardOut()
	StandardErr()
}

// 1.每个进程都被分配了三个文件资源
// 分别是STDIN,STDOUT,STRERR(标准输入,标准输出,错误输出)

func StandardIn() {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func StandardOut() {
	fmt.Println("StandardOut ...")
}

func StandardErr() {
	err := errors.New("error!")
	fmt.Errorf("err ", err)
}
