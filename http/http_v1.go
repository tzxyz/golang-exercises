package main
import (
	"io"
	"net/http"
	"log"
)

func main() {
	http.HandleFunc("/",Hello)
	http.HandleFunc("/hello",Hello)
	http.HandleFunc("/fuck",Fuck)

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func Hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello Beego")
}

func Fuck(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>fuck my life<h1/>")
}