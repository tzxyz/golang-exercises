package main

import "fmt"

func main() {
	a := "im vincent"
	/**
	range是另一种循环形式
	*/
	for i, ch := range a {
		fmt.Println(i, ch)
	}
	fmt.Println("------分割线------")

	b := [...]int{1, 4, 7, 5}

	for a, i := range b {
		fmt.Println(a, i)
	}
}
