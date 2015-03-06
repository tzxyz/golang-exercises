package main

import (
	"io"
	"log"
	"net/http"
)

type HandlerDemo struct{}

func (h *HandlerDemo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "this is a personal handler")
}

func handlertest(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "this is a function handler")
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/structs", &HandlerDemo{})    //处理自定义的handler
	mux.HandleFunc("/functions", handlertest) //处理函数作为handler

	// mux.Handle("/publics", http.StripPrefix("/publics", http.FileServer("/TEST")) 静态文件是怎么处理的

	err := http.ListenAndServe(":9999", mux) //使用自定义的路由处理器
	if err != nil {
		log.Fatal(err)
	}
}
