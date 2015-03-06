package main

import "fmt"

func main() {
	test_for_loop()
	test_goto()
}

func test_for_loop() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("============= test_for_loop =============")
}

func test_goto() {
	i := 0
label:
	if i < 10 {
		fmt.Println(i)
		i++
		goto label
	}
	fmt.Println("============= test_goto =============")
}
