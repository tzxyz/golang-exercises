package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

func main() {
	// GetDemo()
	// PostDemo()
	// PostFormDemo()
	// ClientGetDemo()
	// ClientPostDemo()
	ClientPostFormDemo()
}

func GetDemo() {
	resp, err := http.Get("http://www.huaban.com") //发起一个http get请求
	if err != nil {
		log.Println("err = ", err)
	}
	defer resp.Body.Close()                //关闭资源 q:为什么要写在readall前面
	body, err := ioutil.ReadAll(resp.Body) //从响应体中读取信息
	fmt.Println(string(body))
}

func PostDemo() {
	buf := bytes.NewBuffer([]byte{})
	resp, err := http.Post("http://www.zhihu.com/people/230jkvfan20", "text/html", buf)
	// resp, err := http.Get("http://www.zhihu.com/people/230jkvfan20")
	// resp, err := http.Post("http://www.zhihu.com/register/account", "text/html", buf)
	if err != nil {
		log.Println("err = ", err)
	}
	defer resp.Body.Close() //别忘了关闭resp.Body
	body, err := ioutil.ReadAll(resp.Body)
	content, err := ioutil.ReadAll(buf)
	fmt.Println(string(body))
	fmt.Println(string(content))
	fmt.Println(resp.Status)
	fmt.Println(resp.StatusCode)
}

func PostFormDemo() {
	// 注册知乎成功!!!
	// question: _xsrf是什么东西
	resp, err := http.PostForm("http://www.zhihu.com/register/account",
		url.Values{"_xsrf": {"b2d0bd688909de92bdfbabb04a6c83fe"}, "first_name": {"vfan20"}, "last_name": {"230jk"}, "password": {"@paul0112"}, "email": {"vincent.fml@hotmail.com"}})

	if err != nil {
		log.Println("post form err = ", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	file, err := os.Create("zhihu-post-form.html")
	io.Copy(file, resp.Body) //为什么复制进去只有一点,但是能打印出来(测试,什么都没复制进去)
	fmt.Println(string(body))
	fmt.Println(resp.Header)
	fmt.Println(resp.StatusCode)
}

func ClientGetDemo() {
	// http包中的 DefaultClient就是 &http.Client{}
	// http.Get(url)就是对DefaultClient进行了一层包装
	client := &http.Client{
		Timeout: time.Second * 10, //为什么这里time.Duration
	}
	resp, _ := client.Get("http://www.baidu.com")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("data = ")
	fmt.Println(string(body))
}

func ClientPostDemo() {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	fmt.Println(client)
	// client.Post()方法中的bodyType起什么作用?
	// resp, _ := client.Post("http://zhihu.com", "image/jpg", body);
}

func ClientPostFormDemo() {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	resp, _ := client.PostForm("http://zhihu.com/register/account",
		url.Values{"username": {"zhuonima"}, "password": {"19930112tz"}}) //go的{}除了函数都要在一行?
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	fmt.Println(resp.Status)
}
