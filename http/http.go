package main
import (
	"net/http"
	"fmt"
)

func sayHelloName(w http.ResponseWriter, r *http.Request){	//这个两个参数是怎么转发来的
	r.ParseForm()		//解析请求参数
	fmt.Println("**********************************")
	fmt.Println(r.Form)		
	fmt.Println(r.URL.Path)				//获取URL的路径
	fmt.Println(r.URL.Scheme)			//获取URL的Scheme
	fmt.Println("**********************************")
	for k, v := range r.Form {			//遍历请求参数，将它按照k-v的形式打印
		fmt.Println("key =", k)
		fmt.Println("val =", v)
	}
	fmt.Fprintf(w, "fuck u itheima")
}

func main() {
	http.HandleFunc("/fuck", sayHelloName) 			//绑定路由和访问方法的映射
	err := http.ListenAndServe(":80", nil)			//设置监听的端口
	if err != nil {
		fmt.Println(err)
	}
}