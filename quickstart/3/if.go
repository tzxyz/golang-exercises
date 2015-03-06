package main
import "fmt"
func main() {
	/**
		if后面直接跟表达式
		在if块声明的，如果跟外部冲突，在if块中使用的是在if块声明的变量
		if的大括号必须这样写 不能单独分成一行
	 */
	a := "im vincent"
	if a:="who what why"; len(a)>5 {
		fmt.Println(a)
	}
	fmt.Println(a)
}