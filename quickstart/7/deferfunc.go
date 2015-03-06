package main

import "fmt"

func main() {
	/**
	 *	wsm 是5个5
	 */
	for i := 0; i < 5; i++ {
		defer func() {
			fmt.Printf("%d ", i)
		}()
	}
}
