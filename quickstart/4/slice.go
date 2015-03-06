package main

import "fmt"

func main() {
	/*
		创建切片的第一种方式
		基于数组创建
		可以使用数组的一部分，整个数组，比数组还大的数组
		arr[begin:end] begin默认是0，end默认是arr.length
		range关键字有两个返回值：
		第一个是索引
		第二个是元素的值
	*/
	arr1 := [...]int{1, 4, 5, 3, 6, 8, 9}
	slice1 := arr1[:5] //取前5个元素
	for _, v := range slice1 {
		fmt.Printf("%d ", v)
	}
	fmt.Println()

	/*
		直接创建切片
	*/
	slice2 := make([]int, 5)
	slice3 := make([]int, 5, 10)
	slice4 := []int{1, 2, 3, 4, 5}

	for _, v := range slice2 {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
	for _, v := range slice3 {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
	for _, v := range slice4 {
		fmt.Printf("%d ", v)
	}
	fmt.Println()

	fmt.Println("cap(): ",cap(slice3))
	// 返回的是切片分配内存空间的大小
	fmt.Println("len(): ",len(slice3))
	// 返回的是切片的元素个数
}
