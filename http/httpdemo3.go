package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

var mux map[string]func(w http.ResponseWriter, r *http.Request)

type Handler struct{}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//r.URL.String() 返回的是端口号之后的路由
	log.Println(r.URL.String())
	if h, ok := mux[r.URL.String()]; ok {
		h(w, r)
		return
	}
	io.WriteString(w, "hello svn and git")
}

func main() {
	//实际上http.NewServeMux()就是一个map[string]func(w http.ResponseWriter,r *http.Request)
	//这里不要在声明mux了,如果声明了mux,则不是使用上面那个全局变量
	mux = make(map[string]func(w http.ResponseWriter, r *http.Request))

	//设置路由
	mux["/svn"] = svn
	mux["/git"] = git

	//创建一个http Server
	server := http.Server{
		Addr:        ":8888",          //服务的地址和监听的端口
		Handler:     &Handler{},       //使用那个处理器
		ReadTimeout: time.Second * 10, //超时时间设置
	}

	//监听服务
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func svn(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "svn service is running")
}

func git(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "git service is running")
}
