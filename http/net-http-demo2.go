package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	NewRequestDemo()
}

func NewRequestDemo() {
	//http.NewRequest(method,url,body),body is optional
	// req, err := http.NewRequest("GET", "www.baidu.com", nil)
	u, _ := url.Parse("http://www.baidu.com") //将字符串解析成url
	req := &http.Request{
		Method: "GET",
		URL:    u, //这里是*url.URL类型
		// Header: http.Header{"Content-Type": {`multipart/form-data;boundary="foo123"`}},
		Body: ioutil.NopCloser(new(bytes.Buffer)), //NopClose是什么
	}

	client := &http.Client{}
	resp, _ := client.Do(req) //客户端发送请求
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func ReadRequestDemo() {
	//未知
}
