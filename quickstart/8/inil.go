package main
import "fmt"

type User struct{}

/**
 *	当接口和储存的对象类型都为nil时，接口才为nil
 */
func main() {
	var i interface{}
	fmt.Println(i)
	fmt.Println(i == nil)
	var p User
	i = p
	fmt.Println(i)
	fmt.Println(i == nil)
}