package main
import "fmt"
func main() {
	/**
		&表示取内存的地址
		*通过指针间接访问目标对象
	*/
	a := "im vincent"
	p1 := &a
	var p2 *string = &a
	fmt.Println(p1)
	fmt.Println(p2)
}