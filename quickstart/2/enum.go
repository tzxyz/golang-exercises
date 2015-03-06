package main
import "fmt"

/**
	iota是常量计数器
	从0开始，组中每一定一个常量+1
	每遇到新的一个const关键就会重置为0
*/
const (
	before1 = iota
	a = 1
	b = 2
	c = 3
	after1 = iota
)

const (
	before2 = iota
	x = "im"
	y = "vincent"
	after2 = iota
)

/*
	定义常量组时，如果不提供初始值
	则使用上行表达式
*/
const (
	h, i, j = "who, ", "what, ", "why"
	o, p, q 
)

func main() {
	fmt.Println(before1)
	fmt.Println(before2)
	fmt.Println(after1)
	fmt.Println(after2)
	fmt.Println(h+i+j)
	fmt.Println(o+p+q)
}