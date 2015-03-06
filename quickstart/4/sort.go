package main

import "fmt"

func main() {

	a := [...]int{45, 62, 20, 42, 56}
	length := len(a)
	fmt.Println("排序前")
	fmt.Println(a)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if a[i] > a[j] {
				// go支持多重赋值
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	fmt.Println("排序后")
	fmt.Println(a)
}
