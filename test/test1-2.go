package main

import "fmt"

func main() {
	Fbizz()
}

func Fbizz(){
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			if i%5 ==0 {
				fmt.Println("FizzBuzz======15")
			}
			fmt.Println("Fizz======3")
		} else if i%5 == 0 {
			if i%3 ==0 {
				fmt.Println("FizzBuzz======15")
			}
			fmt.Println("Bizz======5")
		} else {
			fmt.Println(i)
		}
	}
}
