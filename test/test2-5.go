package main

import "fmt"

func Map(f func(int) int, l []int) []int{
	j := make([]int,len(l))	//创建一个切片用于保存调用函数之后的返回值
	for k,v := range l {
		j[k] = f(v)	//将调用函数后的返回值赋值给切片
	}
	return j
}
func main() {
	f := func(i int) int{
		return 5*i
	}
	l := []int{1,3,5,7,9}
	fmt.Println(Map(f,l))
}
