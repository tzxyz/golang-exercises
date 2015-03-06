package main

import "fmt"

func main() {
	/**
	数组是纯粹的值类型
	*/
	arr1 := [...]int{1, 4, 8}
	arr2 := arr1
	fmt.Println(arr1)
	fmt.Println(arr2)

	arr1[0]++
	fmt.Println(arr1)
	fmt.Println(arr2)

	fmt.Println()

	/**
	引用类型需要使用指针
	*/
	arr3 := [...]int{4, 9}
	arr4 := &arr3	
	//&才是取内存地址
	//*代表的是指针类型

	// var arr3 = [2]int{4,9}
	// var arr4 = &arr3
	fmt.Println(arr3)
	fmt.Println(arr4)

	arr3[0]++
	fmt.Println(arr3)
	fmt.Println(arr4)

}
