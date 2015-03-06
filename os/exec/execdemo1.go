package main

import (
	"fmt"
	"os/exec"
)

// spawn
// 衍生新的进程
// 外部进程结束后继续执行本程序
func main() {
	// exec包中提供了执行任意go或非go的程序
	// 并等待外部进程结束后继续执行本程序
	// exec.Command(name sting,args... string)
	// args为传递的参数
	dateCmd := exec.Command("date")

	// 返回命令执行的结果到标准输出上
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	// grepCmd := exec.Command("grep", "sql")
	// grepIn, err := grepCmd.StdinPipe()
	// grepOut, err := grepCmd.StdoutPipe()
	// fmt.Println("> grep")
	// fmt.Println(string(grepOut))

	// 复杂的bash命令使用这样子比较好
	// grep还有一种写法就是使用Pipe(),感觉很麻烦
	grepCmd := exec.Command("bash", "-c", "ls -a | grep sql")
	grepOut, err := grepCmd.Output()
	fmt.Println("> ls -a | grep sql")
	fmt.Println(string(grepOut))

	// 多参数的
	lsCmd := exec.Command("bash", "-c", "ls -a")
	lsOut, err := lsCmd.Output()
	fmt.Println("> ls")
	fmt.Println(string(lsOut))

}
