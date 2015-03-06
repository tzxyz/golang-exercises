package main

import "fmt"

func main() {
	/**
	continue搭配标签使用，直接跳到标签之后
	必须在continue之前定义LABEL吗？
	*/
	LABEL:
		for {
			for i := 0; i < 10; i++ {
				fmt.Println("im vincet")
				if i > 5 {
					fmt.Println(i)
					continue LABEL
				}
			}
		}
	fmt.Println("OK")
}
