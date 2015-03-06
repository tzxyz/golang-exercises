package main

import "fmt"

func num(x int) int{ //斐波那契数列
	if x==1 || x==2 {
		return 1
	}else{
		return num(x-1)+num(x-2)
	}
}

func main() {
	num := num(7)
	fmt.Println(num)
}
