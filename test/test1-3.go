package main

import "fmt"

func main() {
	Show(20)
	Show2(20)
}

func Show(col int) {
	for i := 0; i < col; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("A")
		}
		fmt.Println()
	}
}

func Show2(col int){
	str := "H"
	for i := 0; i < col; i++ {
		fmt.Printf("%s\n",str)
		str += "H"
	}
}
