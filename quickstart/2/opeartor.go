package main
import "fmt"
func main() {
	/**
			a : 0101
			b : 0110
		a & b : 0100	相同为1，不同为0
		a | b : 0111	有一个是1就是1
		a ^ b : 0011	不同为1，相同为0
		a &^ b: 0001	第二个数为1的话，强制将第一个数改为0，在取&
	 */
	a := 5
	b := 6
	fmt.Println(a & b);
	fmt.Println(a | b);
	fmt.Println(a ^ b);
	fmt.Println(a &^ b);

	x := 6
	y := ^6
	/**
		32位编译器如下
		6 : 0000 0000 0000 0100
	   ^6 :	1111 1111 1111 1011
	 */
	fmt.Println(x)
	fmt.Println(y)
}