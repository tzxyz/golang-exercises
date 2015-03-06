package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Getpid())  //获取当前进程的PID
	fmt.Println(os.Getppid()) //获取父进程的PID
}

// 1.PID
// 什么是PID,PID全称 Process ID,用于标示和区分进程的ID,是一个全局唯一的正整数
// PID由系统随机分配

// 2.PPID
// 什么是PPID,Parent Process ID,表示父进程ID

// 3.操作系统的进程都是有PID=1的内核init进程衍生出来的

// 4.go run 命令每次都是启动一个新的go虚拟机来执行进程,所以看出ppid不是1

// 5.使用go build编译文件后在查看父进程id

// question: build之后运行ppid也不是1(mac 环境)
// 猜想: 原因是在iterm2上执行的,iterm2本身也是一个进程
// 而且每次运行的ppid都是相同的,当关闭iterm2时,在重新开启运行ppid发生变化
