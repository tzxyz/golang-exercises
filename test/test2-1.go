package main

import "fmt"

func main() {
	i, j := sort(15, 4)
	fmt.Println(i,j)
}

func sort(i, j int) (a, b int) {
	if i > j {
		i, j = j, i
	}
	return i, j
}
