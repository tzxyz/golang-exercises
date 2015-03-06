package main

import "fmt"

func main() {
	/**
		goto的LABEL可以再goto之后定义？
	*/
	for {
		for i := 0; i < 10; i++ {
			if i > 5 {
				fmt.Println(i)
				goto LABEL
			}
		}
	}
LABEL:
	fmt.Println("OK")
}
