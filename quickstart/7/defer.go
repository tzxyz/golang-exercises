package main
import "fmt"
func main() {
	/**
	 *	defer 像是一个栈
	 *  先进后出
	 */
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ",i)
	}
}