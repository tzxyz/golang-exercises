package main
import (
	"os"
	"fmt"
	"io"
	"io/ioutil"
)
func main() {
	//1.os.Hostname() (hostname string, err error)：返回主机名
	hostname, err := os.Hostname()
	fmt.Println(hostname, err)

	//2.os.Environ() []string ：返回系统环境变量
	envs := os.Environ()
	fmt.Println(envs)

	//3.os.Getenv(key string) string ：返回指定key的环境变量的value
	env := os.Getenv("JAVA_HOME")
	fmt.Println(env)

	//4.os.Setenv(key, value string) ：设置环境变量
	os.Setenv("fuck","u")
	fmt.Println(os.Getenv("fuck"))

	//5.os.Getpagesize() int：返回系统底层内存分页的大小
	size := os.Getpagesize()
	fmt.Println(size)	

	//6.os.Stat(name string) (fi FileInfo, err error)：返回一个文件状态
	fi, err := os.Stat("d://fuck.txt")
	fmt.Println(fi ,err)

	//7.os.IsExist(err error) bool	MARK
	// b := os.IsExist("d://fuck.txt")

	//8.os.Getwd()：返回当前的系统路径
	wd, err := os.Getwd()	//这样居然可以，之前已经声明过err了
	fmt.Println(wd, err)

	//type File相关
	f, err := os.Create("d://fuck.txt")	//os.Create(name string) (f *File, err error) 创建一个文件
	fmt.Println(f, err)

	f, err = os.Open("d://fuck.txt")	//os.Open(name string) (f *File, err error) 打开一个文件，只能用于读取
	filename := f.Name() 		//func (f *File) Name string 返回文件名称
	fmt.Println(filename)
	f.Close()

	f, err = os.Open("d://fuck.txt")
	var b []byte
	len,err := f.Read(b)
	fmt.Println(len,err)
	f.WriteString("what the fuck")		//为什么没写进去
}