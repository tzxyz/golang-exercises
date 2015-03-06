package main

import "fmt"

func printAll(numbers ...int) {
	for _, i := range numbers { //变参会当成一个切片
		fmt.Println(i)
	}
}

func main() {
	printAll(1, 5, 4, 3, 8)
	printAll(4, 7, 5, 9, 1, 12, 41, 45, 4)
}
