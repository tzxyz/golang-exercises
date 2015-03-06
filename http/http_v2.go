package main

import (
	"io"
	"log"
	"net/http"
)

type MyHandle struct {
}

func (h *MyHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {	//处理器必须实现ServeHTTP函数
	io.WriteString(w, "make big new" + r.URL.String())
}

func fuck(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "fuck my life")
}

func main() {
	mux := http.NewServeMux()     //创建一个Server服务
	mux.Handle("/", &MyHandle{})  //注册处理器
	mux.HandleFunc("/fuck", fuck) //注册处理函数

	err := http.ListenAndServe(":80", mux) //监听端口，使用哪个处理器监听
	if err != nil {
		log.Fatal(err)
	}
}
