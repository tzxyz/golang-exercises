package main

import (
	"fmt"
)

func main() {
	c := make(chan bool)
	go func() {
		fmt.Println("concurrecy")
		c <- true
	}()
	<-c
}
