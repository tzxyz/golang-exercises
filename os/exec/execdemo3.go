package main

import (
	"fmt"
	"os/signal"
)

// 系统调用
// 我们要启动一个进程,需要操作系统的系统调用(system call)
// 操作系统的进程运行在内核态 用户的进程运行在用户态

// Unix的思想: 一切皆文件
// 内核给每个访问文件分配了一个文件描述符,在文件新建和打开的时候返回一个非负数
// POSIX定义了三个常量STDIN_FILENO,STDOUT_FILENO,STDERR_FILENO(0,1,2)
func main() {

}
