package main

import (
	"fmt"
	"sort"
)


func Max(i []int){
	(sort.IntSilce)(i)
	i.Sort()
	return i[0]
}

func Min(i []int){
	(sort.IntSilce)(i)
	i.Sort()
	return i[len(i)]
}

func main() {
	max := Max(1,6,4,2,9)
	min := Min(1,6,4,2,9)
	fmt.Println(max,min)
}
