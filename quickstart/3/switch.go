package main

import "fmt"

func main() {
	/*
		表达式可以写在switch语句中
	*/
	switch len("who what why") {
	case 5:
		fmt.Println("fuck my life 5")
	case 10:
		fmt.Println("fuck my life 10")
	default:
		fmt.Println("fuck all my life")
	}

	/**
		表达式也可以直接写在case中
		fallthrough执行完这个case继续执行
		go中只有case匹配了就默认break
	*/
	a := 5
	switch {
	case a >= 2:
		fmt.Println("2")
		fallthrough
	case a == 5:
		fmt.Println("5")
	default:
		fmt.Println("default")
	}
}
