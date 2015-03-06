package main

import (
	"io"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/what", what) //注册路由
	http.HandleFunc("/the", the)
	http.HandleFunc("/fuck", fuck)

	err := http.ListenAndServe(":8080", nil) //监听8080端口,使用默认的handler
	if err != nil {
		log.Fatal(err)
	}
}

func what(w http.ResponseWriter, r *http.Request) { //两个参数,一个是responeWriter,一个是*request
	io.WriteString(w, "today i must be to haohaoxuexi")
}

func the(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "what the fuck")
}

func fuck(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "ffffffffffffuck")
}
