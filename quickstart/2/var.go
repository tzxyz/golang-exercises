package main
import "fmt"
/**
	全部变量声明不能省略var
	但是可是使用并行模式
	所有变量都可以使用并行的方式
*/
var (
	x = 1
	y = 2
	z = 3
)

func main() {
	var a string = "fuck my life"
	b := "im vincent"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}