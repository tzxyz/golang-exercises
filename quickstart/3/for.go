package main
import "fmt"
func main() {
	fmt.Println("第一种for循环")
	for i := 0; i < 10; i++ {
		fmt.Println("im vincent")
	}
	fmt.Println("第二种for循环")

	/**
		类似java中的doWhile
	 */
	a := 0
	for {
		a ++
		if a > 10 {
			break
		}
		fmt.Println("im vincent 2");
	}

	/**
		类似java中的While
	 */
	fmt.Println("第三种for循环")
	a = 0
	for a < 10 {
		fmt.Println("im vincent 3")
		a ++
	}
}