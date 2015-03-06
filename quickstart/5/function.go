package main

import (
	"errors"
	"fmt"
)

/**
	定义函数的方式：
	func 函数名 (参数列表) (返回值类型){
		执行体 ...
	}
 */
func add(a int, b int) (result int, err error) {
	if a < 0 || b < 0 {
		err = errors.New("should be gt 0 number")
		return
	}
	return a + b, nil
}
func main() {
	result, err := add(4, 5)
	fmt.Println(result)

	result, err = add(-4, -5)
	fmt.Println(err)
}
