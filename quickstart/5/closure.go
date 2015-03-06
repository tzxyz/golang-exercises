package main

import "fmt"
/**
	闭包closure：

 */
func main() {
	var i int = 5

	f := func()(func()) {
		var j int = 10
		return func() {
			fmt.Printf("i, j: %d ,%d\n", i, j)
		}
	}()

	f()
	i *= 2
	f()
}
