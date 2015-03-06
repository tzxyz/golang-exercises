package main

import "fmt"

func main() {
	mySlice := make([]int, 5, 10)
	fmt.Println("cap(): ", cap(mySlice))
	fmt.Println("len():", len(mySlice))

	//往切片中追加元素 append是内置函数
	mySlice = append(mySlice, 1, 2, 3)
	for _, v := range mySlice {
		fmt.Printf("%d ", v)
	}

	mySlice2 := []int{9, 8}
	//往切片中追加切片必须在后面加上省略号
	mySlice3 := append(mySlice, mySlice2...)
	for _, v := range mySlice3 {
		fmt.Printf("%d ", v)
	}
}
