package main

import (
	"io"
	"log"
	"net/http"
	// "fmt"
)

// 全局变量mux用于保存URL和映射处理器的关系
var mux map[string]func(http.ResponseWriter,*http.Request)

func main() {
	server := http.Server {		//创建一个服务器
		Addr:":80",				//注意这里有一个冒号					
		Handler:&MyHandle{},
		ReadTimeout:5e9,
	}

	//不小心用了:=就不是使用全局变量
	mux = make(map[string]func(http.ResponseWriter,*http.Request))	//设置URL和处理器的映射关系
	mux["/"] = home
	mux["/fuck"] = fuck

	err := server.ListenAndServe()		//服务器开始监听
	if err!=nil {
		log.Fatal(err)
	}
}

type MyHandle struct {}

func (h *MyHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h, ok := mux[r.URL.String()]; ok {
		h(w,r)
		return 
	}
	io.WriteString(w, "welcome"+r.URL.String())
}

func home(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>welcome my home</h1>")
}

func fuck(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>fuck my life</h1>")
}