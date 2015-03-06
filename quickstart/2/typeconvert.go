package main
import "fmt"
func main() {
	/**
		go不支持隐式转换
		必须可以转换的两种类型才可以转换
	*/
	a := 1.2
	b := int8(a)
	fmt.Println(b)

	c := 65
	d := string(c)
	fmt.Println(d)
}