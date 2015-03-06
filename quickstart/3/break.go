package main

import "fmt"

func main() {
	/**
	break搭配标签使用，直接跳到标签之后
	*/
	LABEL:
		for {
			for i := 0; i < 10; i++ {
				fmt.Println("im vincet")
				break LABEL
			}
		}
	fmt.Println("OK")
}
