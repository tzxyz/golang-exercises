package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

// 执行外部程序,执行外部程序后不返回原程序
func main() {
	// 去环境变量Path里面查找这个命令
	exec.LookPath("ls")
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	argv := []string{"ls", "-a", "-l", "-h"}
	envv := os.Environ()
	err := syscall.Exec(binary, argv, envv)
	if err != nil {
		panic(err)
	}

	// 因为之前执行了一个外部进程，这个命令也没有执行
	b, err := exec.LookPath("top")
	envv := os.Environ()

	// 参数可以为空,环境变量不能为空
	err = syscall.Exec(b, nil, envv)
	if err != nil {
		fmt.Println(err)
	}

	//可以看到return并没有输出,说明程序没有返回
	fmt.Println("return")
}
