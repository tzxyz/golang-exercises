package main

import (
	"fmt"
	"strconv"
)

type Stack struct{
	arr [10]int
	top int
}

func (s *Stack) Push(x int) {
	if s.top<len(s.arr) {
		s.arr[s.top] = x
		s.top++
	}else{
		fmt.Println("超出范围")
	}
}

func (s *Stack) Pop() int{
	if s.top >0 {
		s.top--
		return s.arr[s.top]
	}
    return -1
}

func (s *Stack) String() string {
	str := ``
	for i := 0; i < s.top; i++ {
		str += `[`+strconv.Itoa(i) + `:` +strconv.Itoa(s.arr[i]) + `]`
	}
	return str
}

func main() {
	s := &Stack{}
	s.Push(5)
	s.Push(20)
	fmt.Printf("stack %v\n", s)
	p :=s.Pop()
	fmt.Printf("stack %v, p:%d\n", s,p)
}
